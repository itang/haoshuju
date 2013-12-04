package services

import (
	 "github.com/itang/haoshuju/haoshuju.api/modules"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/api"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/index"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/system"
	_ "github.com/itang/haoshuju/haoshuju.api/modules/tool"
)

type Services interface {
	GetApiApp() modules.ApiApp
	GetClientApps() []modules.ClientApp
}

func GetDefaultServices() Services {
	return &services{}
}

///////////////////////////////////////////////////////////////////////////
// private
type services struct {
}

func (this services) GetApiApp() modules.ApiApp {
	return apiApp
}

func (this services) GetClientApps() []modules.ClientApp {
	haoshujuNet := modules.ClientApp{
		AppBase: modules.AppBase{
			Module: modules.Module{
				Id:      "haoshuju.net",
				Name:    "haoshuju.net",
				Path:    "/",
				Version: "0.0.1",
				Status:  modules.SValid,
			},
			Hostname: "localhost",
			HttpPort: 3000,
			Type:     modules.ARemote,
		},
		AccessKey: modules.UUID(),
		SecretKey: modules.UUID(),
	}
	return []modules.ClientApp{haoshujuNet}
}

var (
	appId  = "haoshuju.api"
	apiApp = modules.ApiApp(
		modules.AppBase{
			Module: modules.Module{
				Id:      appId,
				Name:    appId,
				Version: "0.0.1",
			},
			Hostname: "localhost",
			HttpPort: 5000,
			Type:     modules.ALocal,
			Modules: modules.GetModules(),
		})
)
