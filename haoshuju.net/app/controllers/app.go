package controllers

import (
	"time"

	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
	XRuntime
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Test() revel.Result {
	time.Sleep(time.Second * 2)
	return c.RenderText("test")
}

func (c App) Test1() revel.Result {
	return c.RenderText("test1")
}
