package service

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"laundry/core/jwttoken"
	"laundry/entity"
	"laundry/errtype"
)

type loginSvc struct {}

func NewLoginSvc() *loginSvc {
	return &loginSvc{}
}

func (as *loginSvc) Login(userBody entity.UserBody) (token string, errType *errtype.Error) {
	user, errType := userRepo.FindUserByUsername(userBody.Username)
	if errType != nil {
		if errType.Error == sql.ErrNoRows {
			errType = errtype.LoginError(errType.Error)
			return
		}
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(userBody.Password))
	if err != nil {
		errType = errtype.LoginError(err)
		return
	}

	payload := jwttoken.JWTPayload{
		UserId: user.UserProfileId.Int64,
		Role:   user.Role.String,
	}

	token, err = jwttoken.NewJWTToken().GenerateToken(payload)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}

	return
}