package service

import (
	"database/sql"
	"laundry/entity"
	"laundry/errtype"
)

type userProfileSvc struct {}

func NewUserProfileSvc() *userProfileSvc {
	return &userProfileSvc{}
}

func (*userProfileSvc) CreateUserProfile(db *sql.Tx, userProfile entity.UserProfileRepo) (userProfileId int64, errType *errtype.Error) {
	return userProfileRepo.InsertUserProfile(db, userProfile)
}
