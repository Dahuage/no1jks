package models

// blog 本质上作为一个没有答案的问题
type User struct {
	ID           uint      `json:"id" gorm:"primary_key" sql:"INT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Name         string    `json:"name" gorm:"" sql:"VARCHAR(1024) NOT NULL"`
	Phone        string    `json:"phone" gorm:"" sql:"TEXT"`
	Password     string    `json:"password" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	Avatar       string    `json:"avatar" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	Gender       int       `json:"gender" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	WechatOpenid int       `json:"wechat_openid" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	CreateAt     int `json:"create_at" gorm:""`
	UpdateAt     int `json:"update_at" gorm:""`
	Status       int       `json:"status" gorm:"" sql:"TINYINT NOT NULL DEFAULT 0"`
	Province     string    `json:"province"`
	City         string    `json:"city"`
}
