package handlers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/haoshuju.api/services"
)

func AppInfoHandler(services services.Services, r render.Render) {
	r.JSON(200, services.GetApiApp())
}

func AppInfoPropHandler(services services.Services, params martini.Params) (ret string) {
	apiApp := services.GetApiApp()
	switch params["prop"] {
	case "version":
		ret = apiApp.Version
	default:
		ret = ""
	}
	return
}

func ClientAppsHandler(services services.Services, r render.Render) {
	r.JSON(200, services.GetClientApps())
}
