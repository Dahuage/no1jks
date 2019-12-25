package utils

import "fmt"

type ServiceErr struct {
	Code int
	Msg string
	Display string
}

func (e ServiceErr) Error() string {
	return fmt.Sprintf("%v: %v %v", e.Code, e.Msg, e.Display)
}

var Errs map[string]*ServiceErr

func init(){
	errs := make(map[string]*ServiceErr)
	errs["OK"] = &ServiceErr{0, "success", "ok"}

	// 账户
	errs["PARAM_ERROR"] = &ServiceErr{4000, "badParams", "输入错误"}
	errs["CAPTCHA_ERROR"] = &ServiceErr{4001, "badCaptcha", "验证码错误"}
	errs["PHONE_ERROR"] = &ServiceErr{4002, "badPhone", "手机格式错误"}
	errs["PASSWORD_ERROR"] = &ServiceErr{4003, "passwordError", "密码错误"}
	errs["USER_EXIST"] = &ServiceErr{4004, "userExist", "用户已经存在"}
	errs["NEED_LOGIN"] = &ServiceErr{4004, "needLogin", "请登录"}
	errs["UNPRIVILEGED"] = &ServiceErr{4005, "unprivileged", "无权限"}

	// 服务
	errs["UNKNOWN_ERROR"] = &ServiceErr{5000, "serviceErr", "服务错误请烧熟再试"}


	Errs = errs
}

const DateFormat = "2006-01-02"
const DateTimeFormat = "2006-01-02 15:04:05"