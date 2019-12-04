package dao

import "no1jks/no1jks/models"

func (d *Dao) GetUserByPhone(phone string) (*models.User, bool) {
	var user models.User
	err := d.Mysql.Where("user.phone = ?", phone).Find(&user).Error
	if err != nil {
		return nil, false
	}
	return &user, true
}