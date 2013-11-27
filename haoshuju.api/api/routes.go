package api

import (
	"github.com/codegangsta/martini"
)

func setRoutes(m martini.Router) {
	m.Get("/", IndexHandler)
	m.Get("/appinfo", AppInfoHandler)
	m.Get("/appinfo/:prop", AppInfoPropHandler)
	m.Get("/client-apps", ClientAppsHandler)
  m.Get("/time", ServerTimeHandler)
}
