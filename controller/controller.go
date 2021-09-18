package controller

import (
	"laundry/controller/validation"
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
