package auth

import (
	"encoding/json"
	"net/http"
	
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	//"../../errorcode"
	//"../../libs/apierror"
	validator "gopkg.in/go-playground/validator.v9"
	"fmt"
)

type AuthHandler struct {
	AuthService 	*AuthService          	`inject:""`
	Validate  		*validator.Validate 	`inject:""`
}

func (h *AuthHandler)loginSuccess(w http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"success":  true,
		"token":  data,
	}
	js, err := json.Marshal(resp)
	if err != nil {
		resp := map[string]interface{}{
			"success":  false,
			"error": fmt.Sprintf("%s", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func (h *AuthHandler) failed(w http.ResponseWriter, status int, err error) {
	var errResp map[string]interface{}
	if err != nil {
		errCode := apierror.CodeInternalServerError
		errMsg := err.Error()
		var errData interface{}
		if f, ok := err.(*apierror.APIError); ok {
			errCode = f.Code
			errMsg = f.Message
			errData = f.Data
		}
		errResp = map[string]interface{}{
			"code":    errCode,
			"message": errMsg,
			"data":    errData,
		}
	}
	resp := map[string]interface{}{
		"data":  nil,
		"error": errResp,
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login LoginInput
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(login); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.AuthService.Login(login.Email, login.Password)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.loginSuccess(w, http.StatusOK, res)
}