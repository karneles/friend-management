package member

import (
	"encoding/json"
	"net/http"
	
	"../../errorcode"
	"../../libs/apierror"
	validator "gopkg.in/go-playground/validator.v9"
	"fmt"
)

type MemberHandler struct {
	MemberService 	*MemberService          	`inject:""`
	Validate  		*validator.Validate 		`inject:""`
}

func (h *MemberHandler)updatesSuccess(w http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"success":  true,
		"recipients":  data,
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

func (h *MemberHandler)friendsSuccess(w http.ResponseWriter, status int, data interface{}, dataLen int) {
	resp := map[string]interface{}{
		"success":  true,
		"friends":  data,
		"count": 	dataLen,
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

func (h *MemberHandler)commonSuccess(w http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"success":  data,
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

func (h *MemberHandler) failed(w http.ResponseWriter, status int, err error) {
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


func (h *MemberHandler) StoreMemberHandler(w http.ResponseWriter, r *http.Request) {
	var member MemberInput
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(member); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.StoreMember(member)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) CreateFriendConnectionHandler(w http.ResponseWriter, r *http.Request) {
	var connection FriendsInput
	if err := json.NewDecoder(r.Body).Decode(&connection); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(connection); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.AddFriend(connection)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) ResolveFriendsHandler(w http.ResponseWriter, r *http.Request) {
	var member EmailInput
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(member); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveFriendList(member)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	var res = make([]string, 0)
	for _, member := range members {
		res = append(res, member.Email)
	}
	h.friendsSuccess(w, http.StatusOK, res, len(res))
}

func (h *MemberHandler) ResolveCommonFriendsHandler(w http.ResponseWriter, r *http.Request) {
	var connection FriendsInput
	if err := json.NewDecoder(r.Body).Decode(&connection); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(connection); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveCommonFriendList(connection)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	var res = make([]string, 0)
	for _, member := range members {
		res = append(res, member.Email)
	}
	h.friendsSuccess(w, http.StatusOK, res, len(res))
}

func (h *MemberHandler) SubscribeUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	var updates UpdatesInput
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(updates); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.SubscribeUpdates(updates)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) BlockUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	var updates UpdatesInput
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(updates); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.BlockUpdates(updates)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) ResolveUpdatedMemberHandler(w http.ResponseWriter, r *http.Request) {
	var update SendUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(update); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveUpdatedMember(update)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	var res = make([]string, 0)
	for _, member := range members {
		res = append(res, member.Email)
	}
	h.updatesSuccess(w, http.StatusOK, res)
}