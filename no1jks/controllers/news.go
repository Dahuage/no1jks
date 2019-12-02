package controllers

import (
	"fmt"
)

type NewsHomeController struct {
	baseController
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func (c *NewsHomeController) Get() {
	c.TplName = "no1jks/news.html"
	c.Data["IsLogin"] = false
	c.Data["IsNews"] = "active"
	c.Data["News"] = c.s.GetNewsHomepage(false, 0, nil)
}
