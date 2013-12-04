package tool

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
	m := martini.Classic()
	m.Use(XRuntimeM)
	m.Use(RenderM)

	m.Get("/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)

	return m
}
