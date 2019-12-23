package models

type Answer struct {
	ID         int `json:"id" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	QuestionID int `json:"question_id" gorm:"" sql:"index"`
	UserID     int `json:"user_id" gorm:"" sql:"index"`

	Content         string `json:"content" gorm:"" sql:"TEXT"`
	ViewCount       int    `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int    `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int    `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	DisplayHomepage int    `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int    `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	Score           int    `json:"score" gorm:"" sql:"INT NOT NULL DEFAULT 0"`
	CreateAt        int    `json:"create_at" gorm:""`
	UpdateAt        int    `json:"update_at" gorm:""`
	IsDeleted       int    `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	Conclusion      string `json:"conclusion"`
}
