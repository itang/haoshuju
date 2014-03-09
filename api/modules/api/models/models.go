package models

import (
	"os"
	"strconv"

	"github.com/itang/gotang"
	"github.com/itang/haoshuju/open"
	"github.com/itang/haoshuju/util"
)

type Services interface {
	GetApiApp() open.ApiApp
	GetClientApps() []open.ClientApp
}

func GetDefaultServices() Services {
	return &services{}
}

///////////////////////////////////////////////////////////////////////////
// private
type services struct {
}

func (this services) GetApiApp() open.ApiApp {
	return open.ApiApp(
		open.App{
			Module: open.Module{
				Id:      appId,
				Name:    appId,
				Version: "0.0.1",
			},
			Hostname: "haoshuju.net",
			HttpPort: portFromEnvOr(5000),
			Type:     open.ALocal,
			Modules:  open.GetModules(),
		})
}

func (this services) GetClientApps() []open.ClientApp {
	haoshujuNet := open.ClientApp{
		App: open.App{
			Module: open.Module{
				Id:      "haoshuju.net",
				Name:    "haoshuju.net",
				Path:    "/",
				Version: "0.0.1",
				Status:  open.SValid,
			},
			Hostname: "localhost",
			HttpPort: 3000,
			Type:     open.ARemote,
		},
		AccessKey: util.UUID(),
		SecretKey: util.UUID(),
	}
	return []open.ClientApp{haoshujuNet}
}

var (
	appId = "haoshuju.api"
)

func portFromEnvOr(defaultPort int) int {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return defaultPort
	}
	iport, err := strconv.Atoi(port)
	gotang.AssertNoError(err, "")
	return iport
}
