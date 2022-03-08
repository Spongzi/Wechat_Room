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
	LoginId string `json:"login_id"`
	Tel     int    `json:"tel"`
}

// PramsAddFriend 添加好友模型
type PramsAddFriend struct {
	// 电话号码和账号二选一
	FriendTel     int    `json:"friend_tel"`      // 要添加好友的电话号码--> 查找到用户
	FriendLoginId string `json:"friend_login_id"` // 要添加好友的账号--> 查找用户
	AddFriendMsg  string `json:"add_friend_msg"`  // 添加好友的备注信息
}
