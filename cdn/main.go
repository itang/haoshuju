package main

import (
	"runtime"

	cnd_server "github.com/itang/haoshuju/cdn/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server := cnd_server.NewMartiniServer()
	server.Run()
}
