package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"no1jks/no1jks/service"
	"no1jks/no1jks/utils"
	"strings"
	"time"
)

type AdminUserLoginController struct {
	adminBaseController
}

type AdminUserInfoController struct {
	adminBaseController
}

type AdminUserUploadController struct {
	adminBaseController
}

type adminJsonView struct {
	JsonViewBase
	Token string
}


func (c *AdminUserLoginController) Post() {
	var resp adminJsonView
	var form struct {
		Username string
		Password string
	}
	if parseErr := json.Unmarshal(c.Ctx.Input.RequestBody, &form); parseErr != nil {
		logs.Info("REQUEST body", c.Ctx.Request.Body, parseErr)
		resp.Code = utils.Errs["PARAM_ERROR"].Code
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	admins := "1530024038715010608061admin"
	if !strings.Contains(admins, form.Username) {
		err := utils.Errs["UNPRIVILEGED"]
		resp.Code = err.Code
		resp.Error = *err
		c.Data["json"] = err
		c.ServeJSON()
		return
	}

	u := service.UserVerify{Phone: form.Username, Pass: form.Password}
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
	resp.Token = utils.CreateRandomString(16)
	cache := *(c.s.Dao.Cache)
	cacheErr := cache.Put(resp.Token, user.ID, time.Duration(604800)*time.Second)
	if cacheErr != nil {
		resp.Code = utils.Errs["UNKNOWN_ERROR"].Code
		resp.Error = *utils.Errs["UNKNOWN_ERROR"]
		c.Data["json"] = resp
		c.ServeJSON()
		logs.Debug(cacheErr)
		return
	}
	c.Data["json"] = resp
	c.SetSession("super-jks-admin", user.ID)
	c.ServeJSON()
}

func (c *AdminUserInfoController) Post(){
	var resp adminJsonView
	resp.Code = 0
	resp.Data = map[string]interface{} {
		"roles": []string{"admin"},
		"name":"mm&dh之家",
		"avatar":"",
		"introduction":"we r family.",
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminUserUploadController) Post(){
	var resp adminJsonView
	_, fileHead, _ := c.GetFile("file")
	filePath, visitPath, err := utils.UploadTo(fileHead, "images")
	if err != nil {
		panic(err)
	}
	logs.Info("=======", fileHead.Filename, filePath, visitPath)
	saveErr := c.SaveToFile("file", filePath)
	if saveErr != nil {
		panic(saveErr)
	}
	resp.Data = map[string]string{"file_path": visitPath}
	c.Data["json"] = resp
	c.ServeJSON()
}

