package controllers

import (
	"fmt"
	"time"

	"github.com/robfig/revel"
)

type XRuntime struct {
	*revel.Controller

	Start time.Time
	Cost  time.Duration
	End   time.Time
}

func (c *XRuntime) DoBegin() revel.Result {
	c.Start = time.Now()
	return nil
}

func (c *XRuntime) DoEnd() revel.Result {
	c.End = time.Now()
	c.Cost = c.End.Sub(c.Start)
	c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", c.Cost))
	return nil
}

func init() {
	revel.InterceptMethod((*XRuntime).DoBegin, revel.BEFORE)
	revel.InterceptMethod((*XRuntime).DoEnd, revel.AFTER)
}
