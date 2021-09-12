package service

import (
	"database/sql"
	"laundry/entity"
	"laundry/errtype"
)

type userInfoSvc struct {}

func NewUserInfoSvc() *userInfoSvc {
	return &userInfoSvc{}
}

func (ui *userInfoSvc) GetUserInfo(id int64) (result entity.UserInfo, errType *errtype.Error) {
	userProfile, errType := userProfileRepo.FindUserInfoById(id)
	if errType != nil {
		if errType.Error == sql.ErrNoRows {
			errType = errtype.UnauthorizedError(errType.Error)
			return
		}
		return
	}

	result.UserId = userProfile.Id.Int64
	result.Role = userProfile.Role.String
	result.IsAuthenticated = true

	return
}