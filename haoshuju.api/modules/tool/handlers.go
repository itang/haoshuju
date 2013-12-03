package tool

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/itang/haoshuju/haoshuju.api/modules"
)

type checkResult struct {
	Alive bool `json:"alive"`
}

func CheckHostAliveHandler(params martini.Params, r render.Render) {
	hostaddr := params["hostaddr"]
	if len(hostaddr) > 0 {
		r.JSON(200, checkResult{modules.CheckHostAlive(hostaddr)})
	} else {
		r.JSON(200, checkResult{false})
	}
}
