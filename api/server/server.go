package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/strip"
	"github.com/itang/gotang"
	"github.com/itang/haoshuju/open"

	"github.com/itang/haoshuju/api/modules/api"
	"github.com/itang/haoshuju/api/modules/api/models"
	"github.com/itang/haoshuju/api/modules/index"
	"github.com/itang/haoshuju/api/modules/system"
	"github.com/itang/haoshuju/api/modules/tool"
)

func NewMartiniServer() open.Server {
	app := models.GetDefaultServices().GetApiApp()

	m, ok := index.GetModuleRouter().Handler.(*martini.ClassicMartini)
	gotang.Assert(ok, "index.GetModuleRouter().Handler should be type of *martini.ClassicMartini")

	mountModules(m)

	return &martiniServer{app, m}
}

func mountModules(m *martini.ClassicMartini) {
	mrs := []open.ModuleRouter{api.GetModuleRouter(),
		system.GetModuleRouter(),
		tool.GetModuleRouter(),
		publicModuleRouter(),
	}
	for _, mr := range mrs {
		log.Printf("Loaded module %v", mr.Module.Name)
		m.Any(mr.Module.Path+"/.*", strip.Prefix(mr.Module.Path), mr.Handler.ServeHTTP)
	}
}

func publicModuleRouter() open.ModuleRouter {
	module := open.Module{
		Name: "public",
		Path: "/public",
	}
	m := martini.Classic()
	m.Handlers(martini.Recovery(), martini.Static("public", martini.StaticOptions{SkipLogging: true}))
	return open.ModuleRouter{module, m}
}

type martiniServer struct {
	app open.ApiApp
	m   *martini.ClassicMartini
}

func (server *martiniServer) Name() string {
	return server.app.Name
}

func (server *martiniServer) HttpPort() int {
	return server.app.HttpPort
}

func (server *martiniServer) Run() {
	log.Printf("[martini] listening on port %d, [%s]", server.HttpPort(), server.Name())
	http.ListenAndServe(fmt.Sprintf(":%d", server.HttpPort()), server.m)
}
