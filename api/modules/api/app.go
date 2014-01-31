package api

import (
	"github.com/itang/haoshuju/api/modules/api/routes"
	"github.com/itang/haoshuju/open"
)

func GetModule() open.Module {
	return open.Module{
		Id:   "api",
		Name: "api",
		Path: "/api",
		RestApis: []open.RestApi{
			{Name: "server time",
				Url:    "/api/time",
				Method: "get",
				Status: open.SValid,
			},
			{Name: "alive?",
				Url:    "/api/alive",
				Method: "get",
				Status: open.SValid,
			},
		},
	}
}

func GetModuleRouter() open.ModuleRouter {
	return open.ModuleRouter{GetModule(), routes.Routes()}
}
