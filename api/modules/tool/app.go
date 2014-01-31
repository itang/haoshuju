package tool

import (
	"github.com/itang/haoshuju/open"

	"github.com/itang/haoshuju/api/modules/tool/routes"
)

func GetModule() open.Module {
	return open.Module{
		Id:   "tool",
		Name: "tool",
		Path: "/tool",
		RestApis: []open.RestApi{
			{Name: "check hostaddr alive",
				Url:    "/tool/check-hostaddr-alive/:hostaddr",
				Method: "get",
				Status: open.SValid,
			},
		},
	}
}

func GetModuleRouter() open.ModuleRouter {
	return open.ModuleRouter{GetModule(), routes.Routes()}
}
