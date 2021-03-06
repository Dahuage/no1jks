package controllers

import (
	_ "github.com/astaxie/beego"
	_ "no1jks/no1jks/service"
)

// HomepageController homepage
type HomepageController struct {
	baseController
}

type Banner struct {
	Id     int
	Order  int
	Img    string
	Href   string
	Active string
}

// Get homepage controller
func (c *HomepageController) Get() {
	RenderData := c.s.GetHomeContent(false)
	c.TplName = "no1jks/home.html"
	c.Data["IsHome"] = "active"
	banners := []Banner{
		{1, 1, "/static/imgs/banner2.png", "train", "active"},
		{3, 3, "/static/imgs/banner1.png", "/train", ""},
		{2, 2, "/static/imgs/banner3.png", "/train", ""},
	}
	c.Data["Banners"] = &banners
	c.Data["News"] = (*RenderData)["News"]
	c.Data["Questions"] = (*RenderData)["Questions"]
	c.Data["Books"] = (*RenderData)["Books"]
	c.Data["Blog"] = (*RenderData)["Blog"]
}
