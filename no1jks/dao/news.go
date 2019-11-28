package dao

import (
	"no1jks/no1jks/models"
)

// GetNewsByID as the name said idiot.
func (d *Dao) GetNewsByID(id int) *models.News {
	var news models.News
	db := d.mysql
	db.First(&news, id)
	return &news
}

// GetHomePageNews as the name said idiot.
func (d *Dao) GetHomepageNews(limit uint8) *[]*models.News {
	var news []*models.News
	db := d.mysql
	err := db.Where("display_homepage = ? AND is_deleted = ?",
			 1, models.False).Find(&news)
	if err.Error != nil {
		panic(err.Error)
	}
	return &news
}
