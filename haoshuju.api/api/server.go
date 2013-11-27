package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/martinitang"
)

type Server interface {
	HttpPort() int
	Name() string
	Run()
}

func NewMartiniServer() Server {
	app := GetApiApp()
	m := martini.Classic()

	m.Use(martinitang.XRuntime())
	// render html templates from templates directory
	m.Use(render.Renderer("templates"))
	m.Map(app)
	routes(m)

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
