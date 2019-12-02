package controllers

import "github.com/astaxie/beego/logs"

type QuestionHomeController struct {
	baseController
}

func (c *QuestionHomeController) Get() {
	c.TplName = "no1jks/ask_answer.html"
	c.Data["IsLogin"] = false
	c.Data["IsQuestion"] = "active"
	c.Data["Questions"] = c.s.GetQuestionHomepage(0, false, nil)
	logs.Info("==========", c.Data["Questions"])
}
