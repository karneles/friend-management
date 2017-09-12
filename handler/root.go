package handler

import (
	"github.com/karneles/friend-management/domain/member"
)

// RootHandler should list all the handler that we will use
type RootHandler struct {
	*member.MemberHandler `inject:""`
}