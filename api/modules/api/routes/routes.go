package routes

import (
	"net/http"

	"github.com/codegangsta/martini"

	"github.com/itang/haoshuju/api/modules/api/handlers"
	"github.com/itang/martinitang"
)

func Routes() http.Handler {
	m := martini.Classic()
	m.Handlers(martini.Recovery(), martinitang.XRuntime(), martinitang.DefaultRender())

	m.Get("/time", handlers.ServerTimeHandler)
	m.Get("/alive", handlers.AliveHandler)

	return m
}
