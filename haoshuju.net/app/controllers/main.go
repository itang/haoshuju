package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	gtime "github.com/itang/gotang/time"
	"github.com/itang/haoshuju/haoshuju.net/app"
	"github.com/robfig/revel"
)

type App struct {
	BaseController
}

func (c App) Index() revel.Result {
	c.RenderArgs["now"] = gtime.FormatDefault(time.Now())
	return c.Render()
}

func (c App) Version() revel.Result {
	return c.RenderJson(Message{Level: "success", Data: app.VERSION})
}

type serverTime struct {
	Value int64  `json:"value"`
	Str   string `json:"str"`
}

func (c App) ServerTime() revel.Result {
  var restApiURL = "http://localhost:5000/time"
	var st = serverTime{}
	err := getJSON(restApiURL, &st)
	if err != nil {
    revel.WARN.Printf("get %s, error:%v", restApiURL, err)
		return c.RenderJson(Message{Level: "failture", Data: "unknow"})
	}
	return c.RenderJson(Message{Level: "success", Data: st.Str})
}

func getJSON(url string, ret interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}
	return nil
}
