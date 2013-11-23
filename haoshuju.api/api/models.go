package api

type App struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	HttpPort int    `json:"httpport"`
}

type RestApi struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type AppInfo struct {
	App      App       `json:"app"`
	RestApis []RestApi `json:"restApis"`
}
