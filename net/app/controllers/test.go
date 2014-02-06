package controllers

import (
	"time"

	"github.com/robfig/revel"
)

type Test struct {
	BaseController
}

func (c Test) Test() revel.Result {
	time.Sleep(time.Second * 2)
	return c.RenderText("test")
}

func (c Test) Test1() revel.Result {
	return c.RenderText("test1")
}
