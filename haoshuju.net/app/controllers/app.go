package controllers

import (
	. "github.com/itang/reveltang/controllers"
	"github.com/robfig/revel"
)

type BaseController struct {
	*revel.Controller
	XRuntimeableController
}

type Message struct {
	Code    int
	Level   string
	Message string
	Data    interface{}
}
