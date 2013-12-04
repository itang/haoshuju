package index

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
		Name: "index",
		Path: "/",
	}
}

func GetModuleRouter() open.ModuleRouter {
	m := martini.Classic()
	m.Handlers(martini.Recovery(), XRuntimeM)

	m.Get("/", IndexHandler)

	return open.ModuleRouter{GetModule(), m}
}
