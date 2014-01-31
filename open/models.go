package open

import "fmt"

var modules = map[string]Module{}

func RegistModule(module Module) {
	if _, registed := modules[module.Id]; registed {
		panic(fmt.Sprintf("%v have registed!", module))
	}
	modules[module.Id] = module
}

func GetModules() []Module {
	var ms = make([]Module, 0, len(modules))
	for _, v := range modules {
		ms = append(ms, v)
	}
	return ms
}

func GetModulesMap() map[string]Module {
	return modules
}
