package service

import (
	_ "fmt"
	"github.com/astaxie/beego/logs"
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

func (s *Service) VerifyUser(u *UserVerify) (*models.User, *utils.ServiceErr) {

	p, e := strconv.Atoi(u.Phone)
	if e != nil {
		err := utils.Errs["PHONE_ERROR"]
		return nil, err
	}
	codedPhone := utils.EncodeIntString(p)
	user, ok := s.Dao.GetUserByPhone(codedPhone)
	if !ok {
		logs.Debug("Can't encode ", codedPhone)
		err := utils.Errs["UNKNOWN_ERROR"]
		return nil, err
	}

	saltedPassword, saltErr := utils.EncodeSalt(u.Pass)
	if saltErr != nil {
		logs.Debug("Can't add salt for " + u.Pass)
		err := utils.Errs["UNKNOWN_ERROR"]
		return nil, err
	}
	if saltedPassword == user.Password {
		return user, nil
	}

	err := utils.Errs["PASSWORD_ERROR"]
	return nil, err
}

func (s *Service) CreateUser(u *NewUser) (*models.User, *utils.ServiceErr) {
	p, e := strconv.Atoi(u.Phone)
	if e != nil {
		err := utils.Errs["PHONE_ERROR"]
		return nil, err
	}

	codedPhone := utils.EncodeIntString(p)
	_, ok := s.Dao.GetUserByPhone(codedPhone)
	if ok {
		err := utils.Errs["USER_EXIST"]
		return nil, err
	}

	saltedPassword, saltErr := utils.EncodeSalt(u.Pass)
	reSaltedPassword, reSaltErr := utils.EncodeSalt(u.RePass)
	if saltErr != nil || reSaltErr != nil {
		logs.Debug("Can't add salt for " + u.Pass)
		err := utils.Errs["UNKNOWN_ERROR"]
		return nil, err
	}
	if saltedPassword != reSaltedPassword {
		err := utils.Errs["PASSWORD_ERROR"]
		return nil, err
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
		logs.Debug("Create user err ", nil)
		err := utils.Errs["UNKNOWN_ERROR"]
		return nil, err
	}
	return &user, nil
}
