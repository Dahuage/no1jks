package models

import "time"

type Answer struct {
	ID              uint      `json:"id" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	QuestionID      uint      `json:"question_id" gorm:"" sql:"index"`
	UserID          uint      `json:"user_id" gorm:"" sql:"index"`

	Content         string    `json:"content" gorm:"" sql:"TEXT"`
	ViewCount       int       `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int       `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int       `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	DisplayHomepage int       `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int       `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CreatedAt       time.Time `json:"create_at" gorm:""`
	UpdatedAt       time.Time `json:"update_at" gorm:""`
	IsDeleted       int       `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
}
