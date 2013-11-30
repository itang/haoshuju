package api

import (
	"github.com/codegangsta/martini"
)

func setRoutes(m martini.Router) {
	m.Get("/", IndexHandler)
	m.Get("/system/appinfo", AppInfoHandler)
	m.Get("/system/appinfo/:prop", AppInfoPropHandler)
	m.Get("/system/client-apps", ClientAppsHandler)
	m.Get("/api/time", ServerTimeHandler)
	m.Get("/api/alive", AliveHandler)
	m.Get("/tool/check-hostaddr-alive/:hostaddr", CheckHostAliveHandler)
}
