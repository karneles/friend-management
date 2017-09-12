package handler

import (
	"../domain/member"
)

// RootHandler should list all the handler that we will use
type RootHandler struct {
	*member.MemberHandler `inject:""`
}
