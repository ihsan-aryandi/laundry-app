package controller

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/constant"
	"laundry/core/jwttoken"
	"laundry/errtype"
)

type userInfoCtr struct{}

func NewUserInfoCtr() *userInfoCtr {
	return &userInfoCtr{}
}

func (uic *userInfoCtr) GetUserInfo(ctx *rhaprouter.Context) error {
    return logController(ctx, func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error) {
		cookie, err := ctx.Cookie(constant.AuthorizationCookieName)

		if err != nil {
			errType = errtype.UnauthorizedError(err)
			return
		}

		claims, err := jwttoken.NewJWTToken().ValidateToken(cookie.Value)
		if err != nil {
			errType = errtype.UnauthorizedError(err)
			return
		}

		userInfo, errType := userInfoService.GetUserInfo(claims.UserId)
		if errType != nil {
			return
		}

		res = successRetrieveDataResponse("User Info", userInfo)
		return
 	})
}