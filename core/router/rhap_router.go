package router

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/routes"
	"laundry/warehouse"
)

type rhRouter struct {
	RhRouter *rhaprouter.Router
}

func NewRhRouter() *rhRouter {
	return &rhRouter{RhRouter: rhaprouter.NewRouter()}
}

func (r *rhRouter) Listen(port string) error {
	info := warehouse.Log.NewLogInfo()
	info.Message = fmt.Sprintf("Server started on port %v", port)
	info.StatusCode = 200
	info.Print()

	return r.RhRouter.Listen(port)
}

func (r *rhRouter) SetupRoutes() {
	routes.SetupRoutes(r.RhRouter)
}
