package api

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
