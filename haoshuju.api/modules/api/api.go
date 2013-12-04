package api

import (
	"net/http"

	"github.com/codegangsta/martini"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
)

func init() {
	RegistModule(GetModule())
}

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
	m := martini.Classic()
	m.Use(XRuntimeM)
	m.Use(RenderM)

	m.Get("/time", ServerTimeHandler)
	m.Get("/alive", AliveHandler)
	return m
}
