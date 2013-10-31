package filters

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/revel"
)

func XRuntimeFilter(c *revel.Controller, fc []revel.Filter) {
	if !isStaticResourceRequest(c) {
		start := time.Now()
		doChain(c, fc)
		end := time.Now()
		cost := end.Sub(start)
		c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", cost))
	} else {
		doChain(c, fc)
	}
}

func doChain(c *revel.Controller, fc []revel.Filter) {
	fc[0](c, fc[1:])
}

var statics = []string{".js", ".css", ".gif", ".jpg", ".png", ".map"}

func isStaticResourceRequest(c *revel.Controller) bool {
	for _, v := range statics {
		if strings.HasSuffix(c.Request.RequestURI, v) {
			return true
		}
	}
	return false
}
