package controller

import (
	"errors"
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"net/http"
	"laundry/constant"
	"laundry/entity"
	"laundry/errtype"
	"time"
)

type loginCtr struct{}

func NewLoginCtr() *loginCtr {
	return &loginCtr{}
}

func (ac *loginCtr) Login(ctx *rhaprouter.Context) error {
	return logController(ctx, func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error) {
		var userBody entity.UserBody

		err := ctx.Body(&userBody)
		if err != nil {
			errType = errtype.InvalidRequestBodyError(err)
			return
		}

		if userBody.Email == "" || userBody.Password == "" {
			errType = errtype.LoginError(errors.New("username or password is empty"))
			return
		}

		token, errType := loginService.Login(userBody)
		if errType != nil {
			return
		}

		cookie := &http.Cookie{
			Name:       constant.AuthorizationCookieName,
			Value:      fmt.Sprintf("Bearer %v", token),
			Expires:    time.Now().Add(time.Hour * 24),
			HttpOnly:   true,
		}
		ctx.SetCookie(cookie)

		res = successLoginResponse(nil)
		return
	})
}
