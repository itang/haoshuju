package tool

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/martinitang"
)

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

func GetHandler() http.Handler {
	xruntimeM := martinitang.XRuntime()
	renderM := render.Renderer("templates")

	m := martini.Classic()

	m.Use(xruntimeM)
	m.Use(renderM)
	m.Get("/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)
	return m
}
