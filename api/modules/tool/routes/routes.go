package routes

import (
	"net/http"

	"github.com/codegangsta/martini"

	"github.com/itang/haoshuju/api/modules/tool/handlers"
	"github.com/itang/haoshuju/open/utils"
)

func Routes() http.Handler {
	m := martini.Classic()
	m.Handlers(martini.Logger(), martini.Recovery(), utils.XRuntimeM, utils.RenderM)

	m.Get("/check-hostaddr-alive/:hostaddr", handlers.CheckHostAliveHandler)
	return m
}
