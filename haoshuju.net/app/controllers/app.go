package controllers

import (
	"time"

	gtime "github.com/itang/gotang/time"
	. "github.com/itang/reveltang"
	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
	XRuntimeableController
}

func (c App) Index() revel.Result {
	c.RenderArgs["now"] = gtime.FormatDefault(time.Now())
	return c.Render()
}

func (c App) Test() revel.Result {
	time.Sleep(time.Second * 2)
	return c.RenderText("test")
}

func (c App) Test1() revel.Result {
	return c.RenderText("test1")
}
