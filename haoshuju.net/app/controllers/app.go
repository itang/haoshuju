package controllers

import (
	. "github.com/itang/reveltang/controllers"
	"github.com/robfig/revel"
)

type BaseController struct {
	*revel.Controller
	XRuntimeableController
}
