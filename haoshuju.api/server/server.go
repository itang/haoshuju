package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/strip"
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
	m := index.GetHandler()

	m.Get("/api/.*", strip.Prefix("/api"), api.GetHandler().ServeHTTP)
	m.Get("/tool/.*", strip.Prefix("/tool"), tool.GetHandler().ServeHTTP)
	m.Get("/system/.*", strip.Prefix("/system"), system.GetHandler().ServeHTTP)

	return &martiniServer{app, m}
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
