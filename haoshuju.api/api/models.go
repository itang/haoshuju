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

type AppBase struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Hostname    string    `json:"hostname"`
	HttpPort    int       `json:"httpport"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Description string    `json:"description"`
	Type        AppType   `json:"type"`
}

type ApiApp struct {
	AppBase
	RestApis []RestApi `json:"restApis"`
}

type ApiStatus int

const (
	SValid = iota
	SInValid
)

type RestApi struct {
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Status      ApiStatus `json:"status"`
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
		},
		AccessKey: uuidString(),
		SecretKey: uuidString(),
	}
	return []ClientApp{haoshujuNet}
}

func uuidString() string {
	u4, err := uuid.NewV4()
	gotang.AssertNoError(err)
	return u4.String()
}
