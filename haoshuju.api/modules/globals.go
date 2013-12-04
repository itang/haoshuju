package modules

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/martinitang"
)

var (
	XRuntimeM = martinitang.XRuntime()
	RenderM   = render.Renderer("templates")

	modules = []Module{}
)

func RegistModule(module Module) {
	modules = append(modules, module)
}

func GetModules() []Module {
	return modules
}
