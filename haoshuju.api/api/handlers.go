package api

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info.html", http.StatusFound)
}

func AppInfoHandler(appInfo AppInfo, resp http.ResponseWriter) {
	renderJson(resp, appInfo)
}

func AppInfoPropHandler(appInfo AppInfo, params martini.Params) (ret string) {
	prop := params["prop"]
	switch prop {
	case "version":
		ret = appInfo.App.Version
	default:
		ret = ""
	}

	return
}

//////////////////////////////////////////////////////
// utils
func renderJson(resp http.ResponseWriter, obj interface{}) {
	resp.Header().Add("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(resp)
	encoder.Encode(obj)
}
