package service

import (
	"laundry/errtype"
)

type genderSvc struct {}

func NewGenderSvc() *genderSvc {
	return &genderSvc{}
}

func (*genderSvc) IsGenderExistById(id int64) (bool, *errtype.Error) {
	count, errType := genderRepo.GetGenderCountById(id)
	return count > 0, errType
}