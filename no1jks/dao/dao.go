package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	orm "no1jks/no1jks/utils"
)

// Dao as the name said idiot
type Dao struct {
	mysql *gorm.DB
}

type DataSet struct {
	Page, TotalCount int
}

// New Dao constructor
func New(c *orm.MysqlConf) (d *Dao) {
	d = &Dao{mysql: orm.NewMySQL(c)}
	return d
}

// Close cutoff db connections
func (d *Dao) Close() {
	if d.mysql != nil {
		d.mysql.Close()
	}
}
