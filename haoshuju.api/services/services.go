package services

import (
	"github.com/itang/haoshuju/haoshuju.api/modules"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/api"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/index"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/system"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/tool"
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
	return apiApp
}

func (this services) GetClientApps() []open.ClientApp {
	haoshujuNet := open.ClientApp{
		AppBase: open.AppBase{
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
	appId  = "haoshuju.api"
	apiApp = open.ApiApp(
		open.AppBase{
			Module: open.Module{
				Id:      appId,
				Name:    appId,
				Version: "0.0.1",
			},
			Hostname: "localhost",
			HttpPort: 5000,
			Type:     open.ALocal,
			Modules:  modules.GetModules(),
		})
)
