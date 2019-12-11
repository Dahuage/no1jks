package controllers

import (
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
	page, err := c.GetInt("page")
	if err != nil {
		page = 0
	}
	question, pager := c.s.GetQuestionHomepage(page, false, nil)
	c.Data["Questions"] = question
	c.Data["Pager"] = pager
}


func (c *QuestionDetailController) Get() {
	c.TplName = "no1jks/ask_answer_detail.html"
	c.Data["IsQuestion"] = "active"

	questionId := c.Ctx.Input.Param(":id")
	questionIdInt, err := strconv.Atoi(questionId)
	if err != nil {
		// TODO RETURN 404
	}
	question := c.s.GetQuestionDetail(questionIdInt, nil)
	c.Data["Question"] = question
	breadcrumbs := Breadcrumbs{
		[]struct{Href, Word string}{{"/question", "我有疑问"}},
		(*question).Question.QuestionTitle,
	}
	c.Data["Navigation"] = breadcrumbs
}
