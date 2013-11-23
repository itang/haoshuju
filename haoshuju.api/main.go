package main

import (
	"encoding/json"
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

	m.Get("/", indexHandler)
	m.Get("/appinfo", appInfoHandler)
	m.Get("/appinfo/:prop", appInfoPropHandler)

	run(m)

}

func run(m *martini.ClassicMartini) {
	//m.Run()
	log.Printf("[martini] listening on port %d", appInfo.App.HttpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", appInfo.App.HttpPort), m)
}

//////////////////////////////////////////////////////
// handlers

func indexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info.html", http.StatusFound)
}

func appInfoHandler(resp http.ResponseWriter) {
	encoder := json.NewEncoder(resp)
	encoder.Encode(appInfo)
}

func appInfoPropHandler(params martini.Params) (ret string) {
	prop := params["prop"]
	switch prop {
	case "version":
		ret = appInfo.App.Version
	default:
		ret = ""
	}

	return
}
