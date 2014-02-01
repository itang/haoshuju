package cdn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/itang/gotang"
	"github.com/itang/haoshuju/open"
)

func NewMartiniServer() open.Server {
	m, ok := publicModuleRouter().Handler.(*martini.ClassicMartini)
	gotang.Assert(ok, "public.GetModuleRouter().Handler should be type of *martini.ClassicMartini")

	return &martiniServer{m: m}
}

func publicModuleRouter() open.ModuleRouter {
	module := open.Module{
		Name: "cdn",
		Path: "/",
	}
	m := martini.Classic()
	m.Handlers(martini.Recovery(), martini.Static("public", martini.StaticOptions{SkipLogging: true}))
	return open.ModuleRouter{module, m}
}

type martiniServer struct {
	m *martini.ClassicMartini
}

func (server *martiniServer) Name() string {
	return "cdn"
}

func (server *martiniServer) HttpPort() int {
	return 8081
}

func (server *martiniServer) Run() {
	log.Printf("[martini] listening on port %d, [%s]", server.HttpPort(), server.Name())
	http.ListenAndServe(fmt.Sprintf(":%d", server.HttpPort()), server.m)
}
