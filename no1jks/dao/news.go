package dao

import (
	"no1jks/no1jks/models"
)

// GetNewsByID as the name said idiot.
func (d *Dao) GetNewsByID(id int) (news *models.News) {
	db := d.mysql
	db.First(news, id)
	return
}

