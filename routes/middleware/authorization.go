package middleware

import (
	"database/sql"
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/constant"
	"laundry/controller"
	"laundry/core/jwttoken"
	"laundry/core/response"
	"laundry/entity"
	"laundry/errtype"
	"laundry/repo"
	"laundry/warehouse"
)

func Protect(ctr controller.Controller, roles ...string) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		isMatch, errType := validateUserAndMatchRole(ctx, roles)
		if errType != nil {
			return ctx.JSON(response.Error(errType))
		}

		if !isMatch {
			errType = errtype.ForbiddenError()
			logError(ctx, errType.StatusCode, errType.Message)
			return ctx.JSON(response.Error(errType))
		}

		errLog := warehouse.Log.NewLogError()
		errLog.RequestURI = ctx.Request().RequestURI
		errLog.StatusCode = 500
		errLog.Timestamp = ctx.RequestTime()
		errLog.Method = ctx.Request().Method

		res, errType := ctr(ctx)
		if errType != nil {
			errLog.StatusCode = errType.StatusCode
			errLog.ErrorCode = errType.Code
			errLog.Message = errType.Message
			if errType.Error != nil {
				errLog.Message = fmt.Sprintf("%v (%v)", errLog.Message, errType.Error)
			}
			errLog.Print()

			resp := response.Error(errType)
			resp.Info.Timestamp = ctx.RequestTime()

			return ctx.JSON(resp)
		}
		res.Info.Timestamp = ctx.RequestTime()

		infoLog := warehouse.Log.NewLogInfo()
		infoLog.Timestamp = ctx.RequestTime()
		infoLog.RequestURI = ctx.Request().RequestURI
		infoLog.Method = ctx.Request().Method
		infoLog.StatusCode = res.Info.StatusCode
		infoLog.Message = res.Info.Message
		infoLog.Print()

		return ctx.JSON(res)
	}
}

func getClaims(ctx *rhaprouter.Context) (claims *jwttoken.JWTClaims, errType *errtype.Error) {
	cookie, err := ctx.Cookie(constant.AuthorizationCookieName)
	if err != nil {
		errType = errtype.UnauthorizedError(err)
		return
	}

	claims, err = jwttoken.NewJWTToken().ValidateToken(cookie.Value)
	if err != nil {
		errType = errtype.UnauthorizedError(err)
	}
	return
}

func validateUserAndMatchRole(ctx *rhaprouter.Context, roles []string) (isMatch bool, errType *errtype.Error) {
	claims, errType := getClaims(ctx)
	if errType != nil {
		return
	}

	user, errType := getUserRole(ctx, claims.UserId)
	if errType != nil {
		return
	}

	isMatch = matchRole(user.Role.String, roles)
	return
}

func getUserRole(ctx *rhaprouter.Context, userId int64) (result entity.UserProfileRepo, errType *errtype.Error) {
	userProfileRepo := repo.NewUserProfileRepo()
	result, errType = userProfileRepo.FindUserInfoById(userId)
	if errType != nil {
		if errType.Error == sql.ErrNoRows {
			errType = errtype.ForbiddenError()
			logError(ctx, errType.StatusCode, errType.Message)
		}

		logError(ctx, errType.StatusCode, errType.Message)
	}
	return
}

func matchRole(userRole string, allowedRoles []string) bool {
	if len(allowedRoles) == 0 {
		return true
	}

	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func logError(ctx *rhaprouter.Context, statusCode int, message string) {
	errLog := warehouse.Log.NewLogError()
	errLog.Message = message
	errLog.StatusCode = statusCode
	errLog.Timestamp = ctx.RequestTime()
	errLog.RequestURI = ctx.Request().RequestURI
	errLog.Print()
}
