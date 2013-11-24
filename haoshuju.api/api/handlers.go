package api

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info.html", http.StatusFound)
}

func AppInfoHandler(app ApiApp, resp http.ResponseWriter) {
	renderJson(resp, app)
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

//////////////////////////////////////////////////////
// utils
func renderJson(resp http.ResponseWriter, obj interface{}) {
	resp.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(resp).Encode(obj)
}
