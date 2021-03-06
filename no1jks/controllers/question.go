package controllers

import (
	"no1jks/no1jks/utils"
	"strconv"
)

type QuestionHomeController struct {
	baseController
}

type QuestionDetailController struct {
	baseController
}

type QuestionCreate struct {
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

func (c *QuestionCreate) Post(){
	var resp JsonViewBase
	if c.user == nil {
		resp.Code = utils.Errs["NEED_LOGIN"].Code
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	title := c.GetString("title")
	desc := c.GetString("desc")

	ok, err := c.s.CreateQuestion(c.user, title, desc)
	if ok{
		resp.Code = 200
	} else {
		if err != nil {
			resp.Code = err.Code
			resp.Error = *err
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
