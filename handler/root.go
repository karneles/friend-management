package handler

import (
	"github.com/karneles/friend-management/domain/member"
	"github.com/karneles/friend-management/domain/auth"
	//"../domain/member"
	//"../domain/auth"
)

// RootHandler should list all the handler that we will use
type RootHandler struct {
	*member.MemberHandler `inject:""`
	*auth.AuthHandler `inject:""`
}
