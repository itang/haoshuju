package api

import (
	"time"

	"github.com/itang/gotang"
	"github.com/nu7hatch/gouuid"
)

type AppType string

const (
	ALocal  = AppType("local")
	ARemote = AppType("remote")
)

type Status int

const (
	SValid Status = iota
	SInValid
)

type AppBase struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Hostname    string    `json:"hostname"`
	HttpPort    int       `json:"httpport"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Type        AppType   `json:"type"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
}

type ApiApp struct {
	AppBase
	RestApis []RestApi `json:"restApis"`
}

type RestApi struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type ClientApp struct {
	AppBase
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func GetClientApps() []ClientApp {
	haoshujuNet := ClientApp{
		AppBase: AppBase{
			Id:       "haoshuju.net",
			Name:     "haoshuju.net",
			Version:  "0.0.1",
			Hostname: "localhost",
			HttpPort: 3000,
			Type:     ARemote,
			Status:   SValid,
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
				Url:    "/client-apps",
				Method: "get",
				Status: SValid,
			},
		},
	}
)
