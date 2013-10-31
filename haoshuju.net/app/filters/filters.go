package filters

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/revel"
)

func XRuntimeFilter(c *revel.Controller, filterChain []revel.Filter) {
	if !isStaticResourceRequest(c) {
		start := time.Now()
		doChain(c, filterChain)
		end := time.Now()
		cost := end.Sub(start)
		c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", cost))
	} else {
		doChain(c, filterChain)
	}
}

func doChain(c *revel.Controller, filterChain []revel.Filter) {
	filterChain[0](c, filterChain[1:])
}

var statics = []string{".js", ".css", ".gif", ".jpg", ".png", ".map"}

func isStaticResourceRequest(c *revel.Controller) bool {
	fmt.Println(c.Request.RequestURI)
	for _, v := range statics {
		if strings.HasSuffix(c.Request.RequestURI, v) {
			return true
		}
	}
	return false
}
