package index

import (
	"github.com/itang/haoshuju/api/modules/index/routes"
	"github.com/itang/haoshuju/open"
)

func GetModule() open.Module {
	return open.Module{
		Id:   "index",
		Name: "index",
		Path: "/",
	}
}

func GetModuleRouter() open.ModuleRouter {
	return open.ModuleRouter{GetModule(), routes.Routes()}
}
