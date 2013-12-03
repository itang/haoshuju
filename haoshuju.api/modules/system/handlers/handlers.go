package handlers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/haoshuju.api/services"
	"github.com/itang/martinitang"
)

func AppInfoHandler(r render.Render) {
	r.JSON(200, services.GetApiApp())
}

func AppInfoPropHandler(params martini.Params) (ret string) {
	switch params["prop"] {
	case "version":
		ret = services.GetApiApp().Version
	default:
		ret = ""
	}
	return
}

func ClientAppsHandler(r render.Render) {
	r.JSON(200, services.GetClientApps())
}

func GetHandler() http.Handler {
	xruntimeM := martinitang.XRuntime()
	renderM := render.Renderer("templates")

	m := martini.Classic()

	m.Use(xruntimeM)
	m.Use(renderM)

	m.Get("/appinfo", AppInfoHandler)
	m.Get("/appinfo/:prop", AppInfoPropHandler)
	m.Get("/client-apps", ClientAppsHandler)

	return m
}
