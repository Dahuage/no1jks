package models

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title    string    `json:"title" gorm:""`
	Content  string    `json:"content" gorm:""`
	UserId   int64     `json:"user_id" gorm:""`
}
