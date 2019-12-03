package controllers

type UserLoginController struct {
	baseController
}

type UserSignupController struct {
	baseController
}

type UserTermController struct {
	baseController
}

func (c *UserLoginController)Get() {
	c.TplName = "no1jks/login.html"
}

func (c *UserSignupController)Get() {
	c.TplName = "no1jks/signup.html"
}

func (c *UserTermController)Get() {
	c.TplName = "no1jks/term.html"
}