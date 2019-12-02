package controllers

type GoodHomeController struct {
	baseController
}

func (c *GoodHomeController) Get() {
	c.TplName = "no1jks/goods.html"
	c.Data["IsLogin"] = false
	c.Data["IsMaterial"] = "active"
}
