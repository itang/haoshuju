package handlers

import (
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/info", http.StatusFound)
}

func InfoHandler(r render.Render) {
	r.HTML(200, "modules/index/views/info", nil)
}
