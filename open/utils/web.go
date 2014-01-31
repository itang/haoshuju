package utils

import (
	"fmt"
	"html/template"

	"github.com/codegangsta/martini-contrib/render"

	"github.com/itang/martinitang"
)

var (
	XRuntimeM = martinitang.XRuntime()

	RenderM = render.Renderer(render.Options{
		Delims:     render.Delims{"{[{", "}]}"},
		Directory:  ".",                              // specify what path to load the templates from
		Layout:     "layouts/layout",                 // specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},       // Specify extensions to load for templates
		Funcs:      []template.FuncMap{{"cdn": cdn}}, // Specify helper function maps for templates to access.
	})
)

func cdn(resName string) string {
	const CDN_URL_PREFIX = "http://localhost:3000/public"
	var resourcePath string = ""
	switch {
	case resName == "jquery":
		resourcePath = "/js/vendor/jquery/jquery.min.js"
	case resName == "angular":
		resourcePath = "/js/vendor/angular/angular.min.js"
	default:
		panic(fmt.Sprintf("RESOURCE NOT FOUND! (Cdn(%v)", resName))
	}
	return CDN_URL_PREFIX + resourcePath
}
