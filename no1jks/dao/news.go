package dao

import (
	"github.com/jinzhu/gorm"
	"no1jks/no1jks/models"
)

// GetNewsByID as the name said idiot.
func (d *Dao) GetNewsByID(id int) *models.News {
	var news models.News
	db := d.Mysql
	db.First(&news, id)
	return &news
}

// GetHomePageNews as the name said idiot.
func (d *Dao) GetHomepageNews(limit uint8) *[]*models.News {
	var news []*models.News
	db := d.Mysql
	err := db.Where("display_homepage = ? AND is_deleted = ?",
		1, models.False).Find(&news)
	if err.Error != nil {
		panic(err.Error)
	}
	return &news
}

func listBaseFilter(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = ?", models.False)
}

type NewsHomepageSet struct {
	DataSet
	NewsList []models.News
}

func (d *Dao) GetNewsHomepage(page int, onlyCount bool, filters *map[string]interface{}) interface{} {
	var news NewsHomepageSet
	var totalCount int

	db := d.Mysql.Table("news").Scopes(listBaseFilter)
	db.Count(&totalCount)
	if onlyCount {
		return totalCount
	}
	err := db.Order("news.is_top asc, news.create_at desc").
		Offset(getOffset(page)).
		Limit(models.Limit).
		Scan(&news.NewsList).Error
	if err != nil {
		panic(err)
	}

	news.Page = page
	news.TotalCount = totalCount
	return &news
}
