package controllers

import (
	"time"

	gtime "github.com/itang/gotang/time"
	"github.com/robfig/revel"
)

type App struct {
	BaseController
}

func (c App) Index() revel.Result {
	c.RenderArgs["now"] = gtime.FormatDefault(time.Now())
	return c.Render()
}
