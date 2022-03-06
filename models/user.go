package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     int64  `json:"uuid" gorm:"type:int" `             // uuid
	LoginId  string `json:"login_id" gorm:"type:varchar(64)" ` // 登录id
	Name     string `json:"name" gorm:"type:varchar(64)" `     // 用户名
	Sing     string `json:"sing" gorm:"type:varchar(64)" `     // 签名
	Photo    string `json:"photo" gorm:"type:varchar(64)" `    // 头像
	Address  string `json:"address" gorm:"type:varchar(64)" `  // 地址
	Tel      int    `json:"tel" gorm:"type:int" `              // 电话
	Password string `json:"password" gorm:"type:varchar(64)" ` // 密码
	Account  string `json:"account" gorm:"type:varchar(64)" `  // 这个是微信号
	Email    string `json:"email" gorm:"type:varchar(64)" `    // 邮箱
	Status   string `json:"status" gorm:"type:varchar(64)" `   // 状态
}
