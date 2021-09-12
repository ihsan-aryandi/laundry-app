package validation

import (
	"github.com/bandungrhapsody/rhapvalidator"
	"laundry/entity"
	"laundry/errtype"
)

type registerValidation struct{}

func NewRegisterValidation() *registerValidation {
	return &registerValidation{}
}

func (l *registerValidation) ValidateBody(user *entity.UserBody) (results rhapvalidator.ValidationMessages, errType *errtype.Error) {
	/*
		Request body validation
	*/
	validator := rhapvalidator.NewValidator()
	validator.Validate(user)
	if results = validator.Errors(); results != nil {
		return
	}

	/*
		Username validation
	*/
	usernameCount, errType := userRepo.GetUserCountByUsername(user.Username)
	if errType != nil {
		return
	}
	if usernameCount > 0 {
		errType = errtype.UsernameIsAlreadyRegisteredError()
		return
	}

	/*
		Gender validation
	*/
	genderCount, errType := genderRepo.GetGenderCountById(user.GenderId)
	if errType != nil {
		return
	}
	if genderCount == 0 {
		errType = errtype.DataNotFoundError("Gender")
	}
	return
}
