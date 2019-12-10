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
	question := c.s.GetQuestionDetail(questionIdInt, nil)
	c.Data["Question"] = question
	breadcrumbs := Breadcrumbs{
		[]struct{Href, Word string}{{"/question", "我有疑问"}},
		(*question).Question.QuestionTitle,
	}
	c.Data["Navigation"] = breadcrumbs
}
