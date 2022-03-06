package logic

import (
	"wechat_room/dao"
	"wechat_room/middleware/jwtToken"
	"wechat_room/middleware/uuid"
	"wechat_room/models"
	"wechat_room/utile/code"
)

// 逻辑实现-- User

// Register 注册
func Register(p *models.PramsRegister) (code.MyCode, error) {
	// 首先判断数据中是否存在数据
	if _, myCode, err := dao.CheckUserIsExist(p.LoginId, p.Tel); err != nil {
		return myCode, err
	}
	// 判断用户输入是否合法，并且存入数据库
	u := models.User{
		UUID:     uuid.GenID(),
		LoginId:  p.LoginId,
		Password: p.Password,
		Tel:      p.Tel,
		Name:     p.Name,
	}
	if codeMsg, err := dao.RegisterUserInfo(&u); err != nil {
		return codeMsg, err
	}
	return 0, nil
}

// Login 登录
func Login(p *models.PramsLoginUser) (code.MyCode, string, error) {
	// 判断用户信息是否正确
	u := models.User{
		LoginId:  p.LoginId,
		Password: p.Password,
	}
	err := dao.CheckUserInfo(&u)
	if err != nil {
		return code.LoginIdOrPasswordIsFailed, "", err
	}
	// 获取token
	token, err := jwtToken.GetToken(u.UUID, u.Name)
	return code.SUCCESS, token, nil
}

// CheckUser 查询用户
func CheckUser(p *models.PramsCheckUser) (*models.CheckUser, code.MyCode, error) {
	// 先收集查询的账号信息
	u := models.User{
		Tel:     p.Tel,
		LoginId: p.LoginId,
	}
	user, myCode, err := dao.CheckUser(&u)
	if err != nil {
		return nil, myCode, err
	}
	return user, myCode, nil
}
