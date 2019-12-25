package controllers

import (
	"github.com/astaxie/beego"
	"no1jks/no1jks/models"
	"no1jks/no1jks/service"
	"no1jks/no1jks/utils"
)

type JsonViewBase struct {
	Code  int
	Error utils.ServiceErr
	Data  interface{}
}

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implemented global settings for all other routers.
type baseController struct {
	beego.Controller
	s       *service.Service
	user    *models.User
	isLogin bool
}

type adminBaseController struct {
	baseController
}

// maybe different logic for admin
func (this *adminBaseController) Prepare() {
	// init service
	this.s = service.GetService()
	// auth
	userId := this.GetSession("super-jks-admin")
	if userId != nil {
		currentUser, ok := this.s.Dao.GetUserById(userId.(int))
		if ok {
			this.isLogin = true
			this.user = currentUser
			this.Data["IsLogin"] = true
			this.Data["User"] = currentUser
		} else {
			this.isLogin = false
			this.user = nil
			this.Data["IsLogin"] = false
		}
	} else {
		this.isLogin = false
		this.user = nil
		this.Data["IsLogin"] = false
	}
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

// Prepare implemented Prepare method for baseRouter.
func (this *baseController) Prepare() {

	// init service
	this.s = service.GetService()

	// auth
	userId := this.GetSession("super-jks")
	if userId != nil {
		currentUser, ok := this.s.Dao.GetUserById(userId.(int))
		if ok {
			this.isLogin = true
			this.user = currentUser
			this.Data["IsLogin"] = true
			this.Data["User"] = currentUser
		} else {
			this.isLogin = false
			this.user = nil
			this.Data["IsLogin"] = false
		}
	} else {
		this.isLogin = false
		this.user = nil
		this.Data["IsLogin"] = false
	}
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

type Breadcrumbs struct {
	Parent       []struct{ Href, Word string }
	CurrentTitle string
}
