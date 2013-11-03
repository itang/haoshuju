package controllers

import (
	. "github.com/itang/reveltang"
	"github.com/robfig/revel"
)

type BaseController struct {
	*revel.Controller
	XRuntimeableController
}
