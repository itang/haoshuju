package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/itang/haoshuju/haoshuju.api/api"
	"github.com/itang/martinitang"
)

func main() {
	m := martini.Classic()
	app := api.GetApiApp()

	m.Use(martinitang.XRuntime())
	m.Map(app)
	api.Routes(m)

	run(app, m)
}

func run(app api.ApiApp, handler http.Handler) {
	log.Printf("[martini] listening on port %d, [%s]", app.HttpPort, app.Name)
	http.ListenAndServe(fmt.Sprintf(":%d", app.HttpPort), handler)
}
