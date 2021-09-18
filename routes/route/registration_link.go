package route

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/controller"
)

var (
	registrationLinkController = controller.NewRegistrationLinkCtr()
)

func SetupRegistrationLink(r *rhaprouter.Router) {
	r.POST("/v1/admin/registration-link/laundry", registrationLinkController.CreateLaundryRegistrationLink)
}