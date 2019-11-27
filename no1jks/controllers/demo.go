package controllers

import (
	"github.com/astaxie/beego"
)

// MainController a demo
type MainController struct {
	beego.Controller
}

// Get index page
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "demo.tpl"
}
