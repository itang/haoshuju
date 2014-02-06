package controllers

import (
	"time"

	"github.com/itang/gotang"
	gtime "github.com/itang/gotang/time"
	"github.com/itang/haoshuju/api/modules/api/data"
	"github.com/itang/haoshuju/net/app"
	"github.com/itang/haoshuju/open"
	"github.com/robfig/revel"
)

type App struct {
	BaseController
}

func (c App) Index() revel.Result {
	c.RenderArgs["now"] = gtime.FormatDefault(time.Now())
	return c.Render()
}

func (c App) Version() revel.Result {
	return c.RenderJson(open.RestResponse{Data: app.VERSION})
}

func (c App) ServerTime() revel.Result {
	restApiURL := "http://localhost:5000/api/time"
	resp := data.ServerTimeResponse{}
	if err := gotang.HttpGetAsJSON(restApiURL, &resp); err != nil {
		revel.WARN.Printf("get %s, error:%v", restApiURL, err)
		return c.RenderJson(open.RestResponse{Code: 1, Data: "unknown"})
	}
	return c.RenderJson(resp)
}
