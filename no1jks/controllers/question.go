package controllers

type QuestionHomeController struct {
	baseController
}

func (c *QuestionHomeController) Get() {
	c.TplName = "no1jks/ask_answer.html"
	c.Data["IsQuestion"] = "active"
	c.Data["Questions"] = c.s.GetQuestionHomepage(0, false, nil)
}
