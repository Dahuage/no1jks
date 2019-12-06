package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "no1jks/no1jks/routers"
	"no1jks/no1jks/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	_ = beego.AddFuncMap("human_time", utils.Stamp2Str)
	beego.Run()
}
