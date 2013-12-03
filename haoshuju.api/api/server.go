package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/strip"
	"github.com/itang/martinitang"
)

type Server interface {
	HttpPort() int
	Name() string
	Run()
}

func NewMartiniServer() Server {
	xruntimeM := martinitang.XRuntime()
	renderM := render.Renderer("templates")

	app := GetApiApp()
	m := martini.Classic()

	m.Map(app)
	m.Use(xruntimeM)
	m.Use(renderM)

	m.Get("/", IndexHandler)
	m.Get("/system/appinfo", AppInfoHandler)
	m.Get("/system/appinfo/:prop", AppInfoPropHandler)
	m.Get("/system/client-apps", ClientAppsHandler)

	apiMartini := martini.Classic()
	apiMartini.Map(app)
	apiMartini.Use(xruntimeM)
	apiMartini.Use(renderM)

	apiMartini.Get("/time", ServerTimeHandler)
	apiMartini.Get("/alive", AliveHandler)

	toolMartini := martini.Classic()
	toolMartini.Use(xruntimeM)
	toolMartini.Use(renderM)
	toolMartini.Get("/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)

	m.Get("/api/.*", strip.Prefix("/api"), apiMartini.ServeHTTP)
	m.Get("/tool/.*", strip.Prefix("/tool"), toolMartini.ServeHTTP)

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
