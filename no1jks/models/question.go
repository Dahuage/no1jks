package models

import "time"

// blog 本质上作为一个没有答案的问题
type Question struct {
	ID              uint      `json:"id" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	UserID          uint      `json:"user_id" gorm:"" sql:"INT NOT NULL INDEX"`
	Title           string    `json:"title" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Content         string    `json:"content" gorm:"" sql:"TEXT"`
	ViewCount       int       `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int       `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int       `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	ThumbImg        string    `json:"thumb_img" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	DisplayHomepage int       `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int       `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsBlog          int       `json:"is_blog" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CreatedAt       time.Time `json:"create_at" gorm:""`
	UpdatedAt       time.Time `json:"update_at" gorm:""`
	IsDeleted       int       `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsLocked        int       `json:"is_locked" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
}