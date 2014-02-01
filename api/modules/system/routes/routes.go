package routes

import (
	"net/http"

	"github.com/codegangsta/martini"

	api_models "github.com/itang/haoshuju/api/modules/api/models"
	"github.com/itang/haoshuju/api/modules/system/handlers"
	"github.com/itang/martinitang"
)

func Routes() http.Handler {
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), martinitang.XRuntime(), martinitang.DefaultRender())

	m.Get("/appinfo", servicesInjector, handlers.AppInfoHandler)
	m.Get("/appinfo/:prop", servicesInjector, handlers.AppInfoPropHandler)
	m.Get("/client-apps", servicesInjector, handlers.ClientAppsHandler)

	return m
}

func servicesInjector(context martini.Context) {
	context.MapTo(api_models.GetDefaultServices(), (*api_models.Services)(nil))
}
