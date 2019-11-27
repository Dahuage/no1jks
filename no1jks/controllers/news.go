package controllers

type NewsHomeController struct {
	baseController
}

func (c *NewsHomeController) Get() {
	c.TplName = "no1jks/news.html"
	c.Data["IsLogin"] = false
	c.Data["IsNews"] = "active"
}
