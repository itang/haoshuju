package handlers

import (
	"time"

	"github.com/codegangsta/martini-contrib/render"
	gtime "github.com/itang/gotang/time"
	"github.com/itang/haoshuju/api/modules/api/data"
)

func ServerTimeHandler(r render.Render) {
	now := time.Now()
	r.JSON(200, data.ServerTimeResponse{
		Data: data.ServerTimeResponseData{
			now.Unix(),
			gtime.FormatDefault(now),
		},
	})
}

func AliveHandler(r render.Render) {
	r.JSON(200, data.AliveResponse{
		Data: data.AliveResponseData{true},
	})
}
