package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/strip"
	"github.com/itang/gotang"
	. "github.com/itang/haoshuju/haoshuju.api/modules"

	"github.com/itang/haoshuju/haoshuju.api/modules/api"
	"github.com/itang/haoshuju/haoshuju.api/modules/index"
	system "github.com/itang/haoshuju/haoshuju.api/modules/system/handlers"
	"github.com/itang/haoshuju/haoshuju.api/modules/tool"
	"github.com/itang/haoshuju/haoshuju.api/services"
)

type Server interface {
	HttpPort() int
	Name() string
	Run()
}

func NewMartiniServer() Server {
	app := services.GetDefaultServices().GetApiApp()

	m, ok := index.GetModuleRouter().Handler.(*martini.ClassicMartini)
	gotang.Assert(ok, "should be martini handler")

	mountModules(m)

	return &martiniServer{app, m}
}

func mountModules(m *martini.ClassicMartini) {
	mrs := []ModuleRouter{api.GetModuleRouter(),
		system.GetModuleRouter(),
		tool.GetModuleRouter(),
		publicModuleRouter(),
	}
	for _, mr := range mrs {
		m.Any(mr.Module.Path+"/.*", strip.Prefix(mr.Module.Path), mr.Handler.ServeHTTP)
	}
}

func publicModuleRouter() ModuleRouter {
	module := Module{
		Name: "public",
		Path: "/public",
	}
	m := martini.Classic()
	m.Handlers(martini.Recovery(), martini.Static("public"))
	return ModuleRouter{module, m}
}

type martiniServer struct {
	app ApiApp
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
