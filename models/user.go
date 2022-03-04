package models

type User struct {
	ID       int    `json:"id" db:"id"`
	UUID     int64  `json:"uuid" db:"uuid" `         // uuid
	LoginId  string `json:"login_id" db:"login_id" ` // 登录id
	Name     string `json:"name" db:"name" `         // 用户名
	Sing     string `json:"sing" db:"sing" `         // 签名
	Photo    string `json:"photo" db:"photo" `       // 头像
	Address  string `json:"address" db:"address" `   // 地址
	Tel      int    `json:"tel" db:"tel" `           // 电话
	Password string `json:"password" db:"password" ` // 密码
	Account  string `json:"account" db:"account" `   // 这个是微信号
	Email    string `json:"email" db:"email" `       // 邮箱
	Status   string `json:"status" db:"status" `     // 状态
}
