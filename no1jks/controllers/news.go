package controllers

type NewsHomeController struct {
	baseController
}

func (c *NewsHomeController) Get() {
	c.TplName = "no1jks/news.html"
	c.Data["IsNews"] = "active"
	c.Data["News"] = c.s.GetNewsHomepage(false, 0, nil)
}
