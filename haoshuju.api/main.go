package main

import (
	"github.com/itang/haoshuju/haoshuju.api/api"
)

func main() {
	var server api.Server = api.NewMartiniServer()
	server.Run()
}
