package dao

import (
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"no1jks/no1jks/models"
	orm2 "no1jks/no1jks/utils/orm"
)

// Dao as the name said idiot
type Dao struct {
	Mysql *gorm.DB
	Cache *cache.Cache
	Es    *string
}

type DataSet struct {
	Page, TotalCount int
}

func getOffset(page int) int {
	return page * models.Limit
}

// New Dao constructor
func New(c *orm2.MysqlConf) (d *Dao) {

	// 暂时不写配置了,因为没有别的选择
	str := map[string]string{}
	str["key"] = "jks"
	str["conn"] = ":6379"
	str["dbNum"] = "12"
	bytes, _ := json.Marshal(str)
	redis, err := cache.NewCache("redis", string(bytes))
	if err != nil {
		panic(err)
	}
	d = &Dao{
		Mysql: orm2.NewMySQL(c),
		Cache: &redis,
		Es:    nil,
	}
	return d
}

// Close cutoff db connections
func (d *Dao) Close() {
	if d.Mysql != nil {
		_ = d.Mysql.Close()
	}
}
