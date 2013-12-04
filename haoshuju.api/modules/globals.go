package modules

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/open"
	"github.com/itang/martinitang"
)

var (
	XRuntimeM = martinitang.XRuntime()
	RenderM   = render.Renderer("templates")

	modules = []open.Module{}
)

func RegistModule(module open.Module) {
	modules = append(modules, module)
}

func GetModules() []open.Module {
	return modules
}
