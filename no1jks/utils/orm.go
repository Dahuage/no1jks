package orm

import (
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// MysqlConf As the name said. idot
type MysqlConf struct {
	MysqlDSN         string
	MysqlActive      int
	MysqlIdle        int
	MysqlIdleTimeout int
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	logs.Info(strings.Repeat("%v ", len(v)), v...)
}

// NewMySQL mysql client factory
func NewMySQL(c *MysqlConf) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.MysqlDSN)
	defer db.Close()
	if err != nil {
		logs.Error("cant connect (%s)", c.MysqlDSN)
		logs.Error("db dsn(%s) connect err: %v", c.MysqlDSN, err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(c.MysqlIdle)
	db.DB().SetMaxOpenConns(c.MysqlActive)
	db.DB().SetConnMaxLifetime(time.Duration(c.MysqlIdleTimeout) / time.Second)
	return db
}
