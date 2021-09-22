package controller

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/controller/validation"
	"laundry/core/response"
	"laundry/errtype"
	"laundry/service"
)

/*
	Services, Repositories, Validations
*/
var (
	/*
		Services
	*/
	registerService = service.NewRegisterSvc()
	loginService = service.NewLoginSvc()
	userInfoService = service.NewUserInfoSvc()

	/*
		Repositories
	*/

	/*
		Validations
	*/
	registerValidation = validation.NewRegisterValidation()
)

type Controller func(ctx *rhaprouter.Context) (res *response.APIResponse, errType *errtype.Error)