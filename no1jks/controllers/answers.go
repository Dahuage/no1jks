package controllers

import (
	"no1jks/no1jks/models"
	"no1jks/no1jks/utils"
)

type AnswerCreateController struct {
	baseController
}

type AnswerLikeController struct {
	baseController
}

type AnswerCommentController struct {
	baseController
}

func (c *AnswerCreateController) Post() {
	if c.user == nil {
		c.Redirect("/user/login", 302)
		return
	}
	var resp JsonViewBase
	conclusion := c.GetString("conclusion")
	content := c.GetString("content")
	questionId, _ := c.GetInt("questionId")
	ok, err := c.s.AnswerCreate(c.user, questionId, conclusion, content)
	if ok {
		resp.Code = 200
	} else {
		resp.Code = err.Code
		resp.Error = *err
	}
	c.ServeJSON()
}

func (c *AnswerLikeController) Post() {
	if c.user == nil {
		c.Redirect("/user/login", 302)
		return
	}
	var resp JsonViewBase
	answerId, err := c.GetInt("answerId")
	if err != nil {
		e := utils.Errs["PARAM_ERROR"]
		resp.Code = e.Code
		resp.Error = *e
		c.ServeJSON()
		return
	}
	var answer models.Answer
	db := c.s.Dao.Mysql.First(&answer, answerId)
	if db.Error != nil || answer.ID == 0 {
		e := utils.Errs["PARAM_ERROR"]
		resp.Code = e.Code
		resp.Error = *e
		c.ServeJSON()
		return
	}
	c.s.Dao.Mysql.Model(&answer).Update("like_count", answer.LikeCount + 1)
	resp.Code = 0
	c.ServeJSON()
	return
}
func (c *AnswerCommentController) Post() {}
