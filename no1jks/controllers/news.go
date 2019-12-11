package controllers

import (
	"strconv"
)

type NewsHomeController struct {
	baseController
}

type NewsDetailController struct {
	baseController
}

func (c *NewsHomeController) Get() {
	c.TplName = "no1jks/news.html"
	c.Data["IsNews"] = "active"
	page, err := c.GetInt("page")
	if err != nil {
		page = 0
	}
	news, pager := c.s.GetNewsHomepage(false, page, nil)
	c.Data["News"] = news
	c.Data["Pager"] = pager
}

func (c *NewsDetailController) Get() {
	c.TplName = "no1jks/news_detail.html"
	c.Data["IsNews"] = "active"

	newsId := c.Ctx.Input.Param(":id")
	newsIdInt, err := strconv.Atoi(newsId)
	if err != nil {
		// TODO RETURN 404
		panic("Login")
	}
	news := c.s.GetNewsDetail(newsIdInt, nil)
	c.Data["News"] = news
	breadcrumbs := Breadcrumbs{
		[]struct{Href, Word string}{{"/news", "最新资讯"}},
		(*news).News.Title,
	}
	c.Data["Navigation"] = breadcrumbs
}
