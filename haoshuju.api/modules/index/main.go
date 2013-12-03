package index

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/martinitang"
)

func GetModule() modules.Module {
	return modules.Module{
		Name: "index",
		Path: "/",
	}
}

func GetHandler() *martini.ClassicMartini {
	xruntimeM := martinitang.XRuntime()
	renderM := render.Renderer("templates")

	m := martini.Classic()

	m.Use(xruntimeM)
	m.Use(renderM)

	m.Get("/", IndexHandler)
	return m
}
