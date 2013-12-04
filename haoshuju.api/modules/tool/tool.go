package tool

import (
	"github.com/codegangsta/martini"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
)

func init() {
	RegistModule(GetModule())
}

func GetModule() Module {
	return Module{
		Name: "tool",
		Path: "/tool",
		RestApis: []RestApi{
			{Name: "check hostaddr alive",
				Url:    "/tool/check-hostaddr-alive/:hostaddr",
				Method: "get",
				Status: SValid,
			},
		},
	}
}

func GetModuleRouter() ModuleRouter {
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), XRuntimeM, RenderM)

	m.Get("/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)

	return ModuleRouter{GetModule(), m}
}
