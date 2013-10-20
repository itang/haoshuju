package controllers

import (
	"time"

	"github.com/astaxie/beego"
	gtime "github.com/itang/gotang/time"
)

type MainController struct {
	beego.Controller
}

type beegoInfo struct {
	RunMode  string
	HttpPort int
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["Beego"] = beegoInfo{beego.RunMode, beego.HttpPort}
	this.Data["Now"] = time.Now().Format(gtime.ChinaDefaultDateTimeFull)
	this.TplNames = "index.tpl"
}
