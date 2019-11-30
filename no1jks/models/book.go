package models

import "time"

type Book struct {
	ID              uint   `json:"title" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Sku             string `json:"sku" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Title           string `json:"title" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Desc            string `json:"desc" gorm:"" sql:"TEXT"`
	Recommendation  string `json:"recommendation" gorm:"" sql:"TEXT"`
	Publisher       string `json:"publisher" gorm:"" sql:"VARCHAR(1024)"`
	Author          string `json:"author" gorm:"" sql:"VARCHAR(1024)"`
	DetailInfo      string `json:"detail_info" gorm:"" sql:"TEXT"`
	Price           int    `json:"price" gorm:"" sql:"int"`
	ViewCount       int    `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int    `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int    `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	ThumbImg        string `json:"thumb_img" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	DisplayHomepage int    `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int    `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`

	CreatedAt time.Time `json:"create_at" gorm:""`
	UpdatedAt time.Time `json:"update_at" gorm:""`
	IsDeleted int       `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
}