package handlers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	api_models "github.com/itang/haoshuju/api/modules/api/models"
)

func AppInfoHandler(services api_models.Services, r render.Render) {
	r.JSON(200, services.GetApiApp())
}

func AppInfoPropHandler(services api_models.Services, params martini.Params) (ret string) {
	apiApp := services.GetApiApp()
	switch params["prop"] {
	case "version":
		ret = apiApp.Version
	default:
		ret = ""
	}
	return
}

func ClientAppsHandler(services api_models.Services, r render.Render) {
	r.JSON(200, services.GetClientApps())
}
