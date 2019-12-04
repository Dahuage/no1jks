package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	_ "github.com/astaxie/beego/utils/captcha"
	"no1jks/no1jks/service"
	_ "no1jks/no1jks/service"
)

var cap *captcha.Captcha

func init() {
	store := service.SvrInstance.Dao.Cache
	cap = captcha.NewWithFilter("/captcha/", *store)
}

type UserLoginController struct {
	baseController
}

type UserSignupController struct {
	baseController
}

type UserTermController struct {
	baseController
}

type UserCaptchaController struct {
	baseController
}

func (c *UserLoginController) Get() {
	c.TplName = "no1jks/user_login.html"
	c.Data["Err"] = service.UserVerifyErr{"", "", ""}
}

func (c *UserLoginController) Post() {
	var err service.UserVerifyErr
	logs.Info(c.Ctx.Request.Form)

	verifyCodeErr := cap.VerifyReq(c.Ctx.Request)
	if !verifyCodeErr {
		err.Captcha = "验证码错误"
		c.Data["json"] = err
		c.ServeJSON()
		return
	}
	u := service.UserVerify{}
	if parseErr := c.ParseForm(&u); parseErr != nil {
		err.Captcha = "输入有误请重新输入"
		err.Phone = err.Captcha
		err.Pass = err.Captcha
		c.Data["json"] = err
		c.ServeJSON()
		logs.Debug(parseErr)
		return
	}
	user, e := c.s.VerifyUser(&u)
	if e != nil {
		c.Data["Err"] = e
		c.Data["json"] = e
		c.ServeJSON()
		return
	}
	c.SetSession("super-jks", user.ID)
	c.Redirect("/", 302)
}

func (c *UserSignupController) Get() {
	c.TplName = "no1jks/user_signup.html"
}

func (c *UserSignupController) Post() {
	c.TplName = "no1jks/user_signup.html"
}

func (c *UserTermController) Get() {
	c.TplName = "no1jks/user_term.html"
}
