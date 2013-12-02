package main

import (
	"runtime"

	"github.com/itang/haoshuju/haoshuju.api/api"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var server api.Server = api.NewMartiniServer()
	server.Run()
}
