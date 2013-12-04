package handlers

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/haoshuju/haoshuju.api/services"
)

func GetHandler() http.Handler {
	m := martini.Classic()
	m.Use(modules.XRuntimeM)
	m.Use(modules.RenderM)

	m.Get("/appinfo", servicesInjector, AppInfoHandler)
	m.Get("/appinfo/:prop", servicesInjector, AppInfoPropHandler)
	m.Get("/client-apps", servicesInjector, ClientAppsHandler)

	return m
}

func servicesInjector(context martini.Context) {
	context.MapTo(services.GetDefaultServices(), (*services.Services)(nil))
}
