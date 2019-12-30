package controllers

import (
	"no1jks/no1jks/models"
	"no1jks/no1jks/utils"
	"strconv"
)

type NewsHomeController struct {
	baseController
}

type NewsDetailController struct {
	baseController
}

type NewsLikeController struct {
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
		news.News.Title,
	}
	c.Data["Navigation"] = breadcrumbs
}

func (c *NewsLikeController) Post() {
	if c.user == nil {
		c.Redirect("/user/login", 302)
		return
	}
	var resp JsonViewBase
	NewsId, err := c.GetInt("NewsId")
	if err != nil {
		e := utils.Errs["PARAM_ERROR"]
		resp.Code = e.Code
		resp.Error = *e
		c.ServeJSON()
		return
	}
	var news models.News
	db := c.s.Dao.Mysql.First(&news, NewsId)
	if db.Error != nil || news.ID == 0 {
		e := utils.Errs["PARAM_ERROR"]
		resp.Code = e.Code
		resp.Error = *e
		c.ServeJSON()
		return
	}
	c.s.Dao.Mysql.Model(&news).Update("like_count", news.LikeCount + 1)
	resp.Code = 0
	c.ServeJSON()
	return
}