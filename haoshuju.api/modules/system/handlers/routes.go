package handlers

import (
	"github.com/codegangsta/martini"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/haoshuju/haoshuju.api/modules/system"
	"github.com/itang/haoshuju/haoshuju.api/services"
)

func GetModuleRouter() ModuleRouter {
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), XRuntimeM, RenderM)

	m.Get("/appinfo", servicesInjector, AppInfoHandler)
	m.Get("/appinfo/:prop", servicesInjector, AppInfoPropHandler)
	m.Get("/client-apps", servicesInjector, ClientAppsHandler)

	return ModuleRouter{system.GetModule(), m}
}

func servicesInjector(context martini.Context) {
	context.MapTo(services.GetDefaultServices(), (*services.Services)(nil))
}
