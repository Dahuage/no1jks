package service

import (
	"github.com/astaxie/beego"
	"log"
	"no1jks/no1jks/dao"
	orm2 "no1jks/no1jks/utils/orm"
	"sync"
)

type Service struct {
	dao *dao.Dao
}

type Composite map[string]interface{}

func newService() (s *Service) {
	mysqlActive, err := beego.AppConfig.Int("mysqlActive")
	if err != nil {
		log.Fatal()
	}
	mysqlIdle, err := beego.AppConfig.Int("mysqlIdle")
	if err != nil {
		log.Fatal()
	}
	mysqlIdleTimeout, err := beego.AppConfig.Int("mysqlIdleTimeout")
	if err != nil {
		log.Fatal()
	}
	MySqlConf := &orm2.MysqlConf{
		MysqlDSN:         beego.AppConfig.String("mysqlDsn"),
		MysqlActive:      mysqlActive,
		MysqlIdle:        mysqlIdle,
		MysqlIdleTimeout: mysqlIdleTimeout,
	}
	s = &Service{
		dao: dao.New(MySqlConf),
	}
	return
}

var serviceInstance *Service
var once sync.Once

func GetService() *Service {
	once.Do(func() {
		serviceInstance = newService()
	})
	return serviceInstance
}
