package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/itang/haoshuju/haoshuju.api/api"
)

var (
	appInfo = api.AppInfo{
		api.App{
			Name:     "haoshuju.api",
			Version:  "0.0.1",
			Hostname: "localhost",
			HttpPort: 5000,
		},
		[]api.RestApi{
			{Name: "app info",
				Url:    "/appinfo",
				Method: "get",
			},
			{Name: "app info prop",
				Url:    "/appinfo/:prop",
				Method: "get",
			},
		},
	}
)

func main() {
	m := martini.Classic()

	m.Map(appInfo)
	api.Routes(m)

	run(m)
}

func run(handler http.Handler) {
	//m.Run()
	log.Printf("[martini] listening on port %d", appInfo.App.HttpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", appInfo.App.HttpPort), handler)
}
