package service

import (
	_ "fmt"
	"no1jks/no1jks/models"
	"no1jks/no1jks/utils"
	_ "no1jks/no1jks/utils"
	_ "os"
	"strconv"
)

type UserVerify struct {
	Phone string `form:"phone"`
	Pass string  `form:"pass"`
}

type UserVerifyErr struct {
	Phone string
	Pass string
	Captcha string
}

func (s *Service)VerifyUser(u *UserVerify) (*models.User, *UserVerifyErr)  {
	var err UserVerifyErr

	p, e := strconv.Atoi(u.Phone)
	if e != nil{
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


func (s *Service)CreateUser(u *UserVerify) (*models.User, *UserVerifyErr)  {
	return nil, nil
}
