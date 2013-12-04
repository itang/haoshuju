package index

import (
	"github.com/codegangsta/martini"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
)

func init() {
	RegistModule(GetModule())
}

func GetModule() Module {
	return Module{
		Name: "index",
		Path: "/",
	}
}

func GetModuleRouter() ModuleRouter {
	m := martini.Classic()
	m.Handlers(martini.Recovery(), XRuntimeM)

	m.Get("/", IndexHandler)

	return ModuleRouter{GetModule(), m}
}
