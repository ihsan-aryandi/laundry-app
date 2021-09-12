package controller

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/entity"
	"laundry/errtype"
)

type registerCtr struct{}

func NewRegisterCtr() *registerCtr {
	return &registerCtr{}
}

func (rc *registerCtr) Register(ctx *rhaprouter.Context) error {
	return logController(ctx, func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error) {
		var userBody entity.UserBody

		err := ctx.Body(&userBody)
		if err != nil {
			errType = errtype.InvalidRequestBodyError(err)
			return
		}

		errValidations, errType := registerValidation.ValidateBody(&userBody)
		if errType != nil {
			return
		}
		if errValidations != nil {
			errType = errtype.InvalidFieldsError(errValidations)
			return
		}

		errType = registerService.Register(userBody)
		if errType != nil {
			return
		}

		res = successRegistrationResponse(nil)
		return
	})
}