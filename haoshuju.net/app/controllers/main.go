package controllers

import (
	"time"

	gtime "github.com/itang/gotang/time"
	"github.com/itang/haoshuju/haoshuju.net/app"
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
	return c.RenderJson(Message{Level: "success", Data: app.VERSION})
}

func (c App) ServerTime() revel.Result {
	return c.RenderJson(Message{Level: "success", Data: gtime.FormatDefault(time.Now())})
}
