package routes

import (
	"net/http"

	"github.com/codegangsta/martini"

	"github.com/itang/haoshuju/api/modules/index/handlers"
	"github.com/itang/haoshuju/open/utils"
)

func Routes() http.Handler {
	m := martini.Classic()
	m.Handlers(martini.Recovery(), utils.XRuntimeM, utils.RenderM)

	m.Get("/", handlers.IndexHandler)
	m.Get("/info", handlers.InfoHandler)

	return m
}
