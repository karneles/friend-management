package auth

import (
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	"github.com/karneles/friend-management/domain/member"
	"github.com/karneles/friend-management/libs/common"
	//"../../errorcode"
	//"../../libs/apierror"
	//"../member"
	//"../../libs/common"
)

type AuthService struct {
	MemberService *member.MemberService `inject:""`
}

func (s *AuthService) Login(email string, password string) (string, error) {
	member, err := s.MemberService.ResolveMemberByEmail(email)
	if err != nil {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Invalid username or password.")
		return "", err
	}
	if common.Hash(password) != member.Password {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Invalid username or password.")
		return "", err
	}

	token, err := common.CreateToken(member.ID.String())
	if err != nil {
		return "", err
	}
	return token, nil
}