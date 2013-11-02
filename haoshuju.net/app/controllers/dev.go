package controllers

import (
	"time"

	"github.com/robfig/revel"
)

var serverSince = time.Now().UnixNano() / 1000000.0

type Dev struct {
	*revel.Controller
}

func (c Dev) SourceChanged(clientSince int64) revel.Result {
	if serverSince > clientSince {
		return c.RenderText("true")
	}
	return c.RenderText("false")
}
