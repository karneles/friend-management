/**************
**
** Note: This methods is created as my initiative to propose a better approach 
**       to create the API. In these methods, the user id is added in the request header.
**		 Using this approach, we can create a cleaner and simpler request message and 
**       add better security.
**
**************/

package member

import (
	"encoding/json"
	"net/http"
	
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	"github.com/karneles/friend-management/libs/common"
	//"../../errorcode"
	//"../../libs/apierror"
	//"../../libs/common"
)

func (h *MemberHandler) CreateFriendConnectionHandler2(w http.ResponseWriter, r *http.Request) {
	var friend EmailInput
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.AddFriend2(userId, friend.Email)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) ResolveFriendsHandler2(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveFriendList2(userId)
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

func (h *MemberHandler) ResolveCommonFriendsHandler2(w http.ResponseWriter, r *http.Request) {
	var friend EmailInput
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveCommonFriendList2(userId, friend.Email)
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

func (h *MemberHandler) SubscribeUpdatesHandler2(w http.ResponseWriter, r *http.Request) {
	var friend EmailInput
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.SubscribeUpdates2(userId, friend.Email)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) BlockUpdatesHandler2(w http.ResponseWriter, r *http.Request) {
	var friend EmailInput
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(friend); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	res, err := h.MemberService.BlockUpdates2(userId, friend.Email)
	if err != nil {
		h.failed(w, apierror.GetHTTPStatus(err), err)
		return
	}
	h.commonSuccess(w, http.StatusOK, res)
}

func (h *MemberHandler) ResolveUpdatedMemberHandler2(w http.ResponseWriter, r *http.Request) {
	var text TextInput
	if err := json.NewDecoder(r.Body).Decode(&text); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	if err := h.Validate.Struct(text); err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	token := r.Header.Get("Token")
	userId, err := common.ParseToken(token)
	if err != nil {
		h.failed(w, http.StatusBadRequest, apierror.FromError(errorcode.InvalidRequestData, err))
		return
	}
	members, err := h.MemberService.ResolveUpdatedMember2(userId, text.Text)
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