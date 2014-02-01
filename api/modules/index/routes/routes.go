package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/codegangsta/martini"

	"github.com/itang/haoshuju/api/modules/index/handlers"
	"github.com/itang/martinitang"
)

func Routes() http.Handler {
	m := martini.Classic()

	render := martinitang.MakeCommonRender("modules/index/views/layouts/layout", []template.FuncMap{{"cdn": cdn}}...)
	m.Handlers(martini.Recovery(), martinitang.XRuntime(), render)

	m.Get("/", handlers.IndexHandler)
	m.Get("/info", handlers.InfoHandler)

	return m
}

func cdn(resName string) string {
	const CDN_URL_PREFIX = "http://localhost:8081"
	var resourcePath string = ""
	switch {
	case resName == "jquery":
		resourcePath = "/js/vendor/jquery/jquery.min.js"
	case resName == "angular":
		resourcePath = "/js/vendor/angular/angular.min.js"
	case resName == "bootstrap.css" || resName == "bootstrap":
		resourcePath = "/libs/bootstrap/css/bootstrap.min.css"
	case resName == "bootstrap.js":
		resourcePath = "/libs/bootstrap/js/bootstrap.min.js"
	default:
		panic(fmt.Sprintf("RESOURCE NOT FOUND! (Cdn(%v)", resName))
	}
	return CDN_URL_PREFIX + resourcePath
}
