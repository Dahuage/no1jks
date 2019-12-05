package service

import (
	_ "fmt"
	"no1jks/no1jks/models"
	"no1jks/no1jks/utils"
	_ "no1jks/no1jks/utils"
	_ "os"
	"strconv"
	"time"
)

type UserVerify struct {
	Phone string `form:"phone"`
	Pass  string `form:"pass"`
}

type NewUser struct {
	UserVerify
	Name   string `form:"username"`
	RePass string `form:"repass"`
}

type UserVerifyErr struct {
	Phone   string
	Pass    string
	Captcha string
	Unknown  string
}

func (s *Service) VerifyUser(u *UserVerify) (*models.User, *UserVerifyErr) {
	var err UserVerifyErr

	p, e := strconv.Atoi(u.Phone)
	if e != nil {
		err.Phone = "手机格式错误"
		return nil, &err
	}
	codedPhone := utils.EncodeIntString(p)
	user, ok := s.Dao.GetUserByPhone(codedPhone)
	if !ok {
		err.Phone = "手机号错误"
		return nil, &err
	}

	saltedPassword, saltErr := utils.EncodeSalt(u.Pass)
	if saltErr != nil {
		panic("Can't add salt for " + u.Pass)
	}
	if saltedPassword == user.Password {
		return user, nil
	}
	err.Pass = "密码错误"
	return nil, &err
}

func (s *Service) CreateUser(u *NewUser) (*models.User, *UserVerifyErr) {
	var err UserVerifyErr

	p, e := strconv.Atoi(u.Phone)
	if e != nil {
		err.Phone = "手机格式错误"
		return nil, &err
	}
	codedPhone := utils.EncodeIntString(p)
	_, ok := s.Dao.GetUserByPhone(codedPhone)
	if ok {
		err.Phone = "该手机已经注册"
		return nil, &err
	}

	saltedPassword, saltErr := utils.EncodeSalt(u.Pass)
	reSaltedPassword, reSaltErr := utils.EncodeSalt(u.RePass)
	if saltErr != nil || reSaltErr != nil {
		panic("Can't add salt for " + u.Pass)
	}
	if saltedPassword == reSaltedPassword {
		err.Pass = "两次密码输入不一致"
		return nil, &err
	}

	var user models.User
	user.Password = saltedPassword
	user.Phone = codedPhone
	user.CreateAt = int(time.Now().Unix())
	user.UpdateAt = int(time.Now().Unix())
	if u.Name != ""{
		user.Name = u.Name
	}else{
		user.Name = "努力的" + u.Phone[7:]
	}
	db := s.Dao.Mysql.Create(&user)
	if db.Error != nil {
		err.Unknown = "未知错误稍后重试"
		return nil, &err
	}
	return &user, nil
}
