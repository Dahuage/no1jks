package controllers

import "github.com/astaxie/beego/logs"

type NewsHomeController struct {
	baseController
}

type NewsDetailController struct {
	baseController
}

func (c *NewsHomeController) Get() {
	c.TplName = "no1jks/news.html"
	c.Data["IsNews"] = "active"
	c.Data["News"] = c.s.GetNewsHomepage(false, 0, nil)
}

func (c *NewsDetailController) Get() {
	c.TplName = "no1jks/news_detail.html"
	c.Data["IsNews"] = "active"

	newsId := c.Ctx.Input.Param(":id")
	c.Data["News"] = c.s.GetNewsDetail(newsId)

	logs.Info("==============", newsId)
}
