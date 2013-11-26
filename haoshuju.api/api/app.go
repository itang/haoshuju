package api

import (
	"github.com/itang/gotang"
	"github.com/nu7hatch/gouuid"
)

var (
	appId  = "haoshuju.api"
	apiApp = ApiApp{
		AppBase{
			Id:       appId,
			Name:     appId,
			Version:  "0.0.1",
			Hostname: "localhost",
			HttpPort: 5000,
		},
		[]RestApi{
			{Name: "app info",
				Url:    "/appinfo",
				Method: "get",
			},
			{Name: "app info prop",
				Url:    "/appinfo/:prop",
				Method: "get",
			},
		},
	}
)

func GetApiApp() ApiApp {
	return apiApp
}

func GetClientApps() []ClientApp {
	haoshujuNet := ClientApp{
		AppBase: AppBase{
			Id:       "haoshuju.net",
			Name:     "haoshuju.net",
			Version:  "0.0.1",
			Hostname: "localhost",
			HttpPort: 3000,
		},
		AccessKey: uuidString(),
		SecretKey: uuidString(),
		Type:      string(ClientAppTypeLocal),
	}
	return []ClientApp{haoshujuNet}
}

func uuidString() string {
	u4, err := uuid.NewV4()
	gotang.AssertNoError(err)
	return u4.String()
}
