package modules

import (
	"html/template"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/open"
	"github.com/itang/martinitang"
)

var (
	XRuntimeM = martinitang.XRuntime()

	RenderM = render.Renderer(render.Options{
		Delims:     render.Delims{"{[{", "}]}"},
		Directory:  "templates",                // specify what path to load the templates from
		Layout:     "layout",                   // specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates
		Funcs:      []template.FuncMap{},       // Specify helper function maps for templates to access.
	})

	modules = []open.Module{}
)

func RegistModule(module open.Module) {
	modules = append(modules, module)
}

func GetModules() []open.Module {
	return modules
}
