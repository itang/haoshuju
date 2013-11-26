package api

import (
	"time"
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
}

type ClientApp struct {
	AppBase
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Type      string `json:"type"`
}

type ClientAppType string

const (
	ClientAppTypeLocal = ClientAppType("local")

	ClientAppTypeRemote = ClientAppType("remote")
)
