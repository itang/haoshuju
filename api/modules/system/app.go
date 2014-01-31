package system

import (
	"github.com/itang/haoshuju/api/modules/system/routes"
	"github.com/itang/haoshuju/open"
)

func GetModule() open.Module {
	return open.Module{
		Id:   "system",
		Name: "system",
		Path: "/system",
		RestApis: []open.RestApi{
			{Name: "app info",
				Url:    "/system/appinfo",
				Method: "get",
				Status: open.SValid,
			},
			{Name: "app info prop",
				Url:    "/system/appinfo/:prop",
				Method: "get",
				Status: open.SValid,
			},
			{Name: "client apps",
				Url:    "/system/client-apps",
				Method: "get",
				Status: open.SValid,
			},
		},
	}
}

func GetModuleRouter() open.ModuleRouter {
	return open.ModuleRouter{GetModule(), routes.Routes()}
}
