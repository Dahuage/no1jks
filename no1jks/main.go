package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
	_ "no1jks/no1jks/routers"
	"no1jks/no1jks/utils"
)

func main() {
	// TODO MOVE ALL THIS SHIT TO CONFIG
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	logs.SetLogger(logs.AdapterFile,`{"filename":"/var/log/no1jks/service.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	_ = beego.AddFuncMap("human_time", utils.Stamp2Str)
	beego.Run()
}
