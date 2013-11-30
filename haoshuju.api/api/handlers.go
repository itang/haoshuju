package api

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	gtime "github.com/itang/gotang/time"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info.html", http.StatusFound)
}

func AppInfoHandler(app ApiApp, r render.Render) {
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

func ServerTimeHandler(r render.Render) {
	now := time.Now()
	r.JSON(200, struct {
		Value int64  `json:"value"`
		Str   string `json:"str"`
	}{
		now.Unix(),
		gtime.FormatDefault(now),
	})
}

func AliveHandler() string {
	return "true"
}

type checkResult struct {
	Alive bool `json:"alive"`
}

func CheckHostAliveHandler(params martini.Params, r render.Render) {
	hostaddr := params["hostaddr"]
	if len(hostaddr) > 0 {
		r.JSON(200, checkResult{CheckHostAlive(hostaddr)})
	} else {
		r.JSON(200, checkResult{false})
	}
}
