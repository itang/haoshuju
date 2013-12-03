package main

import (
	"runtime"

	"github.com/itang/haoshuju/haoshuju.api/app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var server app.Server = app.NewMartiniServer()
	server.Run()
}
