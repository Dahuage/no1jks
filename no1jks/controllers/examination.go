package controllers

type ExaminationHomeController struct {
	baseController
}

func (c *ExaminationHomeController) Get() {
	c.TplName = "no1jks/examination.html"
	c.Data["IsExamination"] = "active"
}
