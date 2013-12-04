package tool

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
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), XRuntimeM, RenderM)

	m.Get("/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)

	return open.ModuleRouter{GetModule(), m}
}
