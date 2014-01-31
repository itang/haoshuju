package routes

import (
	"net/http"

	"github.com/codegangsta/martini"

	"github.com/itang/haoshuju/api/modules/system/handlers"
	"github.com/itang/haoshuju/api/services"
	"github.com/itang/haoshuju/open/utils"
)

func Routes() http.Handler {
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), utils.XRuntimeM, utils.RenderM)

	m.Get("/appinfo", servicesInjector, handlers.AppInfoHandler)
	m.Get("/appinfo/:prop", servicesInjector, handlers.AppInfoPropHandler)
	m.Get("/client-apps", servicesInjector, handlers.ClientAppsHandler)
	return m
}

func servicesInjector(context martini.Context) {
	context.MapTo(services.GetDefaultServices(), (*services.Services)(nil))
}
