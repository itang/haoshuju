package system

import (
	. "github.com/itang/haoshuju/haoshuju.api/modules"
)

func GetModule() Module {
	return Module{
		Name: "system",
		Path: "/system",
		RestApis: []RestApi{
			{Name: "app info",
				Url:    "/system/appinfo",
				Method: "get",
				Status: SValid,
			},
			{Name: "app info prop",
				Url:    "/system/appinfo/:prop",
				Method: "get",
				Status: SValid,
			},
			{Name: "client apps",
				Url:    "/system/client-apps",
				Method: "get",
				Status: SValid,
			},
		},
	}
}
