package api

import (
	"github.com/codegangsta/martini"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/haoshuju/open"
)

func init() {
	RegistModule(GetModule())
}

func GetModule() open.Module {
	return open.Module{
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
	m := martini.Classic()
	m.Handlers(martini.Recovery(), XRuntimeM, RenderM)

	m.Get("/time", ServerTimeHandler)
	m.Get("/alive", AliveHandler)

	return open.ModuleRouter{GetModule(), m}
}
