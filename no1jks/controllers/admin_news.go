package controllers

import (
	"encoding/json"
	"no1jks/no1jks/models"
	"github.com/astaxie/beego/logs"
	"no1jks/no1jks/utils"
)

type AdminNewsController struct {
	baseController
}

type AdminNewsDetailController struct {
	baseController
}

type AdminNewsCreateController struct{
	baseController
}


func (c *AdminNewsController) Get(){
	var resp adminJsonView
	page, err := c.GetInt("page")
	if err != nil {
		page = 0
	} else {
		page -= 1
	}
	newsList, pager := c.s.GetNewsHomepage(false, page, nil)
	resp.Code = 0
	resp.Data = map[string]interface{}{
		"News": newsList,
		"Page": pager,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminNewsDetailController) Get(){
	var resp adminJsonView
	newsId, parseErr := c.GetInt("id")
	if parseErr != nil {

	}
	news := c.s.GetNewsDetail(newsId, nil)
	if news == nil {
		logs.Info("cant find id ==", newsId)
		return
	}
	var data = map[string]interface{}  {
		"status": "draft",
		"title": news.News.Title,
		"content": news.News.Content, // 文章内容
		"content_short": news.News.Brief,  // 文章摘要
		// source_uri: '', // 文章外链
		//image_uri: '', // 文章图片
		//display_time: undefined, // 前台展示时间
		"id": news.News.ID,
		//platforms: ['a-platform'],
		//comment_disabled: false,
		"display_homepage": false,
		"importance": 0,
		"source": "原创",
		"order": 9999,
	}

	resp.Code = 0
	resp.Data = data
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminNewsCreateController) Post(){
	var news models.News
	var resp adminJsonView
	if parseErr := json.Unmarshal(c.Ctx.Input.RequestBody, &news); parseErr != nil {
		logs.Info("REQUEST body", c.Ctx.Request.Body, parseErr)
		resp.Code = utils.Errs["PARAM_ERROR"].Code
		resp.Error.Display = utils.Errs["PARAM_ERROR"].Display
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	db := c.s.Dao.Mysql.Create(&news)
	if err := db.Error;  err != nil {
		logs.Error("Create question err", err, news)
		resp.Code = utils.Errs["PARAM_ERROR"].Code
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	resp.Code = 0
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminNewsCreateController) Put(){
	var news models.News
	var resp adminJsonView
	if parseErr := json.Unmarshal(c.Ctx.Input.RequestBody, &news); parseErr != nil {
		logs.Info("REQUEST body", c.Ctx.Request.Body, parseErr)
		resp.Code = utils.Errs["PARAM_ERROR"].Code
		resp.Error.Display = utils.Errs["PARAM_ERROR"].Display
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	newsId := news.ID
	var savedNews models.News
	//db := c.s.Dao.Mysql.First(&savedNews, newsId)
	//if queryErr := db.Error; queryErr != nil {
	//	panic(queryErr)
	//}
	err := c.s.Dao.Mysql.Model(&savedNews).
				  Where("news.id = ?", newsId).
		          Update(news).Error
	if err != nil {
		panic(err)
	}
	resp.Code = 0
	c.Data["json"] = resp
	c.ServeJSON()
}
