package route

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/controller"
)

var (
	registerController = controller.NewRegisterCtr()
	loginController = controller.NewLoginCtr()
	logoutController = controller.NewLogOutCtr()
	userInfoController = controller.NewUserInfoCtr()
)

func SetupAuth(r *rhaprouter.Router) {
	r.POST("/auth/register", registerController.Register)
	r.POST("/auth/login", loginController.Login)
	r.POST("/auth/logout", logoutController.Logout)
	r.GET("/auth/info", userInfoController.GetUserInfo)
}
