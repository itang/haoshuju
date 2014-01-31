package main

import (
	"runtime"

	"github.com/itang/haoshuju/api/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server := server.NewMartiniServer()
	server.Run()
}
