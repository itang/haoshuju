package api

import (
	"github.com/codegangsta/martini"
)

func routes(m martini.Router) {
	m.Get("/", IndexHandler)
	m.Get("/appinfo", AppInfoHandler)
	m.Get("/appinfo/:prop", AppInfoPropHandler)
	m.Get("/clientapps", ClientAppsHandler)
}
