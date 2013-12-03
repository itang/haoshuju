package api

import (
	"time"

	"github.com/codegangsta/martini-contrib/render"
	gtime "github.com/itang/gotang/time"
)

func ServerTimeHandler(r render.Render) {
	now := time.Now()
	r.JSON(200, struct {
		Value int64  `json:"value"`
		Str   string `json:"str"`
	}{
		now.Unix(),
		gtime.FormatDefault(now),
	})
}

func AliveHandler() string {
	return "true"
}
