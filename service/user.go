package service

import (
	"laundry/errtype"
	"laundry/repo"
)

type userSvc struct {}

var (
	userRepo = repo.NewUserRepo()
)

func NewUserSvc() *userSvc {
	return &userSvc{}
}

func (u *userSvc) IsUserExistByUsername(username string) (bool, *errtype.Error) {
    count, errType := userRepo.GetUserCountByUsername(username)
	return count > 0, errType
}