package data

import (
	"github.com/itang/haoshuju/open"
)

type ServerTimeResponseData struct {
	Value int64  `json:"value"`
	Str   string `json:"str"`
}

type ServerTimeResponse struct {
	open.RestResponseBase
	Data ServerTimeResponseData `json:"data"`
}

type AliveResponseData struct {
	Alive bool `json:"alive"`
}

type AliveResponse struct {
	open.RestResponseBase
	Data AliveResponseData `json:"data"`
}
