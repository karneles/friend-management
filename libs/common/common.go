package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	//"../../errorcode"
	//"../../libs/apierror"
	
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

const (
	SigningKey = "JNEWQndiwq2"
)

type FriendManagementClaims struct {
	UserID	string	`json:"UserID"`
	jwt.StandardClaims
}

func Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateToken(id string) (string, error) {
	mySigningKey := []byte(SigningKey)

	claims := FriendManagementClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 5).Unix(),
			Issuer:    "fm",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &FriendManagementClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	
	if token.Valid {
		claims := token.Claims.(*FriendManagementClaims)
		userId, err := uuid.FromString(claims.UserID)
		if err != nil {
			return uuid.NewV4(), nil
		}
		return userId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			err = apierror.WithMessage(errorcode.ValidationError, "Invalid token")
			return uuid.NewV4(), err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			err = apierror.WithMessage(errorcode.ValidationError, "Token expire")
			return uuid.NewV4(), err
		} else {
			err = apierror.WithMessage(errorcode.ValidationError, "Invalid token")
			return uuid.NewV4(), err
		}
	} else {
		err = apierror.WithMessage(errorcode.ValidationError, "Invalid token")
		return uuid.NewV4(), err
	}

	return uuid.NewV4(), nil
}