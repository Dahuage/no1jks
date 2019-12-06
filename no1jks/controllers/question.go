package controllers

import (
	"github.com/astaxie/beego/logs"
	"strconv"
)

type QuestionHomeController struct {
	baseController
}

type QuestionDetailController struct {
	baseController
}

func (c *QuestionHomeController) Get() {
	c.TplName = "no1jks/ask_answer.html"
	c.Data["IsQuestion"] = "active"
	c.Data["Questions"] = c.s.GetQuestionHomepage(0, false, nil)
}

func (c *QuestionDetailController) Get() {
	c.TplName = "no1jks/ask_answer_detail.html"
	c.Data["IsQuestion"] = "active"

	questionId := c.Ctx.Input.Param(":id")
	questionIdInt, err := strconv.Atoi(questionId)
	if err != nil {
		// TODO RETURN 404
	}
	c.Data["Question"] = c.s.GetQuestionDetail(questionIdInt, nil)
	logs.Info("=======================????", c.Data["Question"])
}
