package controller

import (
	"errors"
	"github.com/bandungrhapsody/rhaprouter"
	"github.com/bandungrhapsody/rhapvalidator"
	"laundry/entity"
	"laundry/errtype"
)

type registrationLinkCtr struct{}

func NewRegistrationLinkCtr() *registrationLinkCtr {
	return &registrationLinkCtr{}
}

func (rl *registrationLinkCtr) CreateLaundryRegistrationLink(ctx *rhaprouter.Context) error {
    return logController(ctx, func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error) {
		var body entity.RegistrationLinkBody

    	err := ctx.Body(&body)
    	if err != nil {
    		errType = errtype.InvalidRequestBodyError(err)
    		return
		}

		/*
			Is email valid ?
		*/
		if !rhapvalidator.IsEmail(body.Email) {
			errType = errtype.InvalidFormatError("email", errors.New("invalid email format"))
			return
		}



    	return
	})
}