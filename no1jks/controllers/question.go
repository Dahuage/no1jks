package controllers

type QuestionHomeController struct {
	baseController
}

func (c *QuestionHomeController) Get() {
	c.TplName = "no1jks/ask_answer.html"
	c.Data["IsLogin"] = false
	c.Data["IsQuestion"] = "active"
}
