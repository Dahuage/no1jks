package models

type News struct {
	ID              uint   `json:"id" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Title           string `json:"title" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Tags            string
	Brief           string
	Content         string `json:"content" gorm:"" sql:"TEXT"`
	ViewCount       int    `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int    `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int    `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	ThumbImg        string `json:"thumb_img" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	DisplayHomepage int    `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int    `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	SourceName      string `json:"source_name" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	SourceHref      string `json:"source_href" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	CreateAt        int    `json:"create_at" gorm:""`
	UpdateAt        int    `json:"update_at" gorm:""`
	IsDeleted       int    `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
}

type NewsComment struct {
	ID              uint   `json:"title" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Title           string `json:"title" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Content         string `json:"content" gorm:"" sql:"TEXT"`
	ViewCount       int    `json:"view_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	LikeCount       int    `json:"like_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CommentCount    int    `json:"comment_count" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	ThumbImg        string `json:"thumb_img" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	DisplayHomepage int    `json:"display_homepage" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	IsTop           int    `json:"is_top" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	SourceName      string `json:"source_name" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	SourceHref      string `json:"source_herf" gorm:"" sql:"VARCHAR(1024) NOT NULL DEFAULT ''"`
	CreatedAt       int    `json:"create_at" gorm:""`
	UpdatedAt       int    `json:"update_at" gorm:""`
	IsDeleted       int    `json:"is_deleted" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
}
