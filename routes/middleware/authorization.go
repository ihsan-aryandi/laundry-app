package middleware

import (
	"database/sql"
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

func Allow(ctr controller.Controller, roles ...string) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		isMatch, errType := validateUserAndMatchRole(ctx, roles)
		if errType != nil {
			return ctx.JSON(response.Error(errType))
		}

		if !isMatch {
			errType = errtype.ForbiddenError()
			logError(ctx, errType)
			return ctx.JSON(response.Error(errType))
		}

		res, errType := ctr(ctx)
		if errType != nil {
			logError(ctx, errType)

			resp := response.Error(errType)
			resp.Info.Timestamp = ctx.RequestTime()
			return ctx.JSON(resp)
		}
		res.Info.Timestamp = ctx.RequestTime()

		infoLog := warehouse.Log.NewLogInfo()
		infoLog.Timestamp = ctx.RequestTime()
		infoLog.RequestURI = ctx.Request().RequestURI
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
	if isMatch {
		ctx.AddContext(constant.UserPayloadContextName, claims)
	}
	return
}

func getUserRole(ctx *rhaprouter.Context, userId int64) (result entity.UserProfileRepo, errType *errtype.Error) {
	userProfileRepo := repo.NewUserProfileRepo()
	result, errType = userProfileRepo.FindUserInfoById(userId)
	if errType != nil {
		if errType.Error == sql.ErrNoRows {
			errType = errtype.ForbiddenError()
			logError(ctx, errType)
		}

		logError(ctx, errType)
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

func logError(ctx *rhaprouter.Context, errType *errtype.Error) {
	errLog := warehouse.Log.NewLogError()
	errLog.Timestamp = ctx.RequestTime()
	errLog.RequestURI = ctx.Request().RequestURI
	errLog.Method = ctx.Request().Method

	if errType == nil {
		errLog.Print()
		return
	}

	errLog.ErrorCode = errType.Code
	errLog.Message = errType.Message
	errLog.StatusCode = errType.StatusCode
	if errType.Error != nil {
		errLog.CausedBy = errType.Error.Error()
	}
	errLog.Print()
}
