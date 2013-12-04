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

func GetHandler() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(XRuntimeM)
	m.Use(RenderM)

	m.Get("/", IndexHandler)

	return m
}
