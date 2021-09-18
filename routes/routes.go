package routes

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/routes/middleware"
	"laundry/routes/route"
)

func SetupRoutes(r *rhaprouter.Router) {
	/*
		Routes Setup
	*/
	route.SetupAuth(r)
	route.SetupRegistrationLink(r)
	/*
		Global Middleware Setup
	*/
	r.Use(middleware.Recover, middleware.Logger, middleware.HandleCORS)
}
