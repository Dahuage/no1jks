package controllers

import (
	"github.com/astaxie/beego"
	"no1jks/no1jks/service"
	"time"
)

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implemented global settings for all other routers.
type baseController struct {
	beego.Controller
	//i18n.Locale
	s *service.Service
	//user    models.User
	//isLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (this *baseController) Prepare() {

	// page start time
	this.s = service.GetService()
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	//this.Data["AppDescription"] = utils.AppDescription
	//this.Data["AppKeywords"] = utils.AppKeywords
	//this.Data["AppName"] = utils.AppName
	//this.Data["AppVer"] = utils.AppVer
	//this.Data["AppUrl"] = utils.AppUrl
	//this.Data["AppLogo"] = utils.AppLogo
	//this.Data["AvatarURL"] = utils.AvatarURL
	//this.Data["IsProMode"] = utils.IsProMode

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
