package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	_ "github.com/astaxie/beego/utils/captcha"
	"no1jks/no1jks/service"
	_ "no1jks/no1jks/service"
	"no1jks/no1jks/utils"
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

type UserContactController struct {
	baseController
}

func (c *UserLoginController) Get() {
	c.TplName = "no1jks/user_login.html"
	c.Data["Login"] = "active"
}

func (c *UserLoginController) Post() {
	logs.Info(c.Ctx.Request.Form)
	var resp JsonViewBase

	verifyCodeErr := cap.VerifyReq(c.Ctx.Request)
	if !verifyCodeErr {
		err := utils.Errs["CAPTCHA_ERROR"]
		resp.Code = err.Code
		resp.Error = *err
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	u := service.UserVerify{}
	if parseErr := c.ParseForm(&u); parseErr != nil {
		err := utils.Errs["PARAM_ERROR"]
		resp.Code = err.Code
		resp.Error = *err
		c.Data["json"] = err
		c.ServeJSON()
		logs.Debug(parseErr)
		return
	}
	user, verifyErr := c.s.VerifyUser(&u)
	if verifyErr != nil {
		resp.Code = verifyErr.Code
		resp.Error = *verifyErr
		c.Data["json"] = resp
		c.ServeJSON()
		logs.Debug(verifyErr)
		return
	}
	resp.Code = 0
	c.Data["json"] = resp
	c.SetSession("super-jks", user.ID)
	c.ServeJSON()
}

func (c *UserSignupController) Get() {
	c.TplName = "no1jks/user_signup.html"
	c.Data["Login"] = "active"
}

func (c *UserSignupController) Post() {
	var resp JsonViewBase
	logs.Info(c.Ctx.Request.Form)

	VerifyOk := cap.VerifyReq(c.Ctx.Request)
	if !VerifyOk {
		err := utils.Errs["CAPTCHA_ERROR"]
		resp.Code = err.Code
		resp.Error = *err
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	u := service.NewUser{}
	if parseErr := c.ParseForm(&u); parseErr != nil {
		err := utils.Errs["PARAM_ERROR"]
		resp.Code = err.Code
		resp.Error = *err
		c.Data["json"] = err
		c.ServeJSON()
		logs.Debug(parseErr)
		return
	}

	user, createErr := c.s.CreateUser(&u)
	if createErr != nil {
		resp.Code = createErr.Code
		resp.Error = *createErr
		c.Data["json"] = resp
		c.ServeJSON()
		logs.Debug(createErr)
		return
	}

	resp.Code = 0
	c.Data["json"] = resp
	c.SetSession("super-jks", user.ID)
	c.ServeJSON()
	return
}

func (c *UserTermController) Get() {
	c.TplName = "no1jks/user_term.html"
}

func (c *UserContactController) Get() {
	c.TplName = "no1jks/contact.html"
	c.Data["IsContact"] = "active"
}


