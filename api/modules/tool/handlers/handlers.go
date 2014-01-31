package handlers

import (
	"net/url"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"github.com/itang/haoshuju/open"
	"github.com/itang/haoshuju/util"
)

type checkResult struct {
	Alive bool `json:"alive"`
}

func CheckHostAliveHandler(params martini.Params, r render.Render) {
	var (
		hostaddr = params["hostaddr"]
		alive    = false
		code     = 0
		message  = ""
		err      = isValidHostAddr(hostaddr)
	)

	if err == nil {
		alive = util.CheckHostAlive(hostaddr)
	} else {
		code = 1
		message = err.Error()
	}

	r.JSON(200, open.RestResponse{
		Code:    code,
		Message: message,
		Data:    checkResult{alive},
	})
}

func isValidHostAddr(s string) error {
	//TODO
	_, err := url.Parse(s)
	return err
}
