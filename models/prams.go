package models

type PramsRegister struct {
	LoginId    string `json:"login_id" binding:"required,gte=6,lte=15"`
	Password   string `json:"password" binding:"required,gte=8,lte=16"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Tel        int    `json:"tel" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Account    string `json:"account" binding:"required"`
}

// PramsLoginUser 登录验证模型
type PramsLoginUser struct {
	LoginId  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PramsCheckUser 查询用户
type PramsCheckUser struct {
	Tel     int    `json:"tel"`
	Account string `json:"account"`
}
