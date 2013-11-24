package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/itang/haoshuju/haoshuju.api/api"
)

func main() {
	m := martini.Classic()

	m.Map(api.GetApiApp())
	api.Routes(m)

	run(m)
}

func run(handler http.Handler) {
	httpPort := api.GetApiApp().HttpPort
	log.Printf("[martini] listening on port %d", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), handler)
}
