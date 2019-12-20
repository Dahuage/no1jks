package controllers

import "github.com/astaxie/beego/logs"

type GoodHomeController struct {
	baseController
}

func (c *GoodHomeController) Get() {
	c.TplName = "no1jks/goods.html"
	c.Data["IsMaterial"] = "active"
	c.Data["Books"] = c.s.GetBooksHomepage(0, false, nil)
	logs.Info("=================", c.Data["Books"])
}
