package system

import (
	. "github.com/itang/haoshuju/haoshuju.api/modules"
	"github.com/itang/haoshuju/open"
)

func init() {
	RegistModule(GetModule())
}

func GetModule() open.Module {
	return open.Module{
		Name: "system",
		Path: "/system",
		RestApis: []open.RestApi{
			{Name: "app info",
				Url:    "/system/appinfo",
				Method: "get",
				Status: open.SValid,
			},
			{Name: "app info prop",
				Url:    "/system/appinfo/:prop",
				Method: "get",
				Status: open.SValid,
			},
			{Name: "client apps",
				Url:    "/system/client-apps",
				Method: "get",
				Status: open.SValid,
			},
		},
	}
}
