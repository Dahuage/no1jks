package models

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title             string `json:"title" gorm:""`
	Content           string `json:"content" gorm:""`
	ViewCount         int    `json:"view_count" gorm:""`
	LikeCount         int    `json:"like_count" gorm:""`
	CommentCount      int    `json:"commet_count" gorm:""`
	ThumbImg          string `json:"thumb_img" gorm:""`
	DisplayHomepage   int    `json:"display_homepage" gorm:""`
	IsTop             int    `json:"is_top" gorm:""`
	SourceName        string `json:"source_name" gorm:""`
	SourceHref        string `json:"source_herf" gorm:""`
}
