package controller

import (
	"github.com/bandungrhapsody/rhaprouter"
	"net/http"
	"laundry/constant"
	"laundry/errtype"
	"time"
)

type logoutCtr struct{}

func NewLogOutCtr() *logoutCtr {
	return &logoutCtr{}
}

func (lc *logoutCtr) Logout(ctx *rhaprouter.Context) error {
    return logController(ctx, func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error) {
		ctx.SetCookie(&http.Cookie{
			Name:       constant.AuthorizationCookieName,
			Expires:    time.Now().Add(-1),
			HttpOnly:   true,
		})

		res = successLogoutResponse(nil)
		return
	})
}