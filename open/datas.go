package open

import (
	"net/http"
	"time"
)

// App type
type AppType string

const (
	ALocal  = AppType("local")
	ARemote = AppType("remote")
)

type Status int

const (
	SValid   Status = iota // 0
	SInValid               // 1
)

type RestApi struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type Module struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Version     string    `json:"version"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
	RestApis    []RestApi `json:"restApis"`
}

type ModuleRouter struct {
	Module  Module
	Handler http.Handler
}

type App struct {
	Module
	Hostname string   `json:"hostname"`
	HttpPort int      `json:"httpport"`
	Type     AppType  `json:"type"`
	Modules  []Module `json:"modules"`
}

type ApiApp App

type ClientApp struct {
	App
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type Server interface {
	HttpPort() int
	Name() string
	Run()
}

type RestResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Data interface {} `json:"data"`
}

type RestResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
