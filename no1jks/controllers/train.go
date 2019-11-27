package controllers

type TrainHomeController struct {
	baseController
}

func (c *TrainHomeController) Get() {
	c.TplName = "no1jks/train.html"
	c.Data["IsLogin"] = false
	c.Data["IsTrain"] = "active"
}
