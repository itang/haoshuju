package api

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info.html", http.StatusFound)
}

func AppInfoHandler(app ApiApp, r render.Render, resp http.ResponseWriter) {
	r.JSON(200, app)
}

func AppInfoPropHandler(app ApiApp, params martini.Params) (ret string) {
	switch params["prop"] {
	case "version":
		ret = app.Version
	default:
		ret = ""
	}
	return
}

func ClientAppsHandler(r render.Render) {
	r.JSON(200, GetClientApps())
}
