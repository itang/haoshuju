package services

import (
	"github.com/itang/gotang"
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/haoshuju/haoshuju.api/modules/api"
	"github.com/itang/haoshuju/haoshuju.api/modules/index"
	"github.com/itang/haoshuju/haoshuju.api/modules/system"
	"github.com/itang/haoshuju/haoshuju.api/modules/tool"
	"github.com/nu7hatch/gouuid"
)

func GetClientApps() []ClientApp {
	haoshujuNet := ClientApp{
		AppBase: AppBase{
			Module: Module{
				Id:      "haoshuju.net",
				Name:    "haoshuju.net",
				Path:    "/",
				Version: "0.0.1",
				Status:  SValid,
			},
			Hostname: "localhost",
			HttpPort: 3000,
			Type:     ARemote,
		},
		AccessKey: uuidString(),
		SecretKey: uuidString(),
	}
	return []ClientApp{haoshujuNet}
}

func GetApiApp() ApiApp {
	return apiApp
}

func uuidString() string {
	u4, err := uuid.NewV4()
	gotang.AssertNoError(err)
	return u4.String()
}

var (
	appId  = "haoshuju.api"
	apiApp = ApiApp(
		AppBase{
			Module: Module{
				Id:      appId,
				Name:    appId,
				Version: "0.0.1",
			},
			Hostname: "localhost",
			HttpPort: 5000,
			Type:     ALocal,
			Modules: []Module{
				api.GetModule(),
				index.GetModule(),
				system.GetModule(),
				tool.GetModule(),
			},
		})
)
