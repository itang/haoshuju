package api

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/martinitang"
)

func GetModule() Module {
	return Module{
		Name: "api",
		Path: "/api",
		RestApis: []RestApi{
			{Name: "server time",
				Url:    "/api/time",
				Method: "get",
				Status: SValid,
			},
			{Name: "alive?",
				Url:    "/api/alive",
				Method: "get",
				Status: SValid,
			},
		},
	}
}

func GetHandler() http.Handler {
	xruntimeM := martinitang.XRuntime()
	renderM := render.Renderer("templates")

	apiMartini := martini.Classic()
	apiMartini.Use(xruntimeM)
	apiMartini.Use(renderM)

	apiMartini.Get("/time", ServerTimeHandler)
	apiMartini.Get("/alive", AliveHandler)
	return apiMartini
}
