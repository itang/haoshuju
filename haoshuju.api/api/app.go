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
			Type:     ALocal,
		},
		[]RestApi{
			{Name: "app info",
				Url:    "/appinfo",
				Method: "get",
				Status: SValid,
			},
			{Name: "app info prop",
				Url:    "/appinfo/:prop",
				Method: "get",
				Status: SValid,
			},
			{Name: "client apps",
				Url:    "/clientapps",
				Method: "get",
				Status: SValid,
			},
		},
	}
)

func GetApiApp() ApiApp {
	return apiApp
}
