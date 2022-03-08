package dao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"wechat_room/models"
	"wechat_room/utile/code"
	"wechat_room/utile/errMsg"
)

// CheckUserIsExist 检查用户是否存在
func CheckUserIsExist(loginId string, tel int) (IsExist bool, myCode code.MyCode, err error) {
	fmt.Println("CheckUserIsExist", loginId, "电话号码是", tel)
	var user models.User
	err = DB.Debug().
		Select("login_id, tel").
		Where("login_id = ?", loginId).Or("tel = ?", tel).
		Find(&user).
		Error
	if err != nil {
		zap.L().Error("检查用户是否存在错误", zap.Error(err))
		myCode = code.CheckFailed
		IsExist = false
		return IsExist, myCode, errMsg.CheckFailed
	}
	fmt.Println("用户登录", user.LoginId, "用户手机号验证", user.Tel)
	if user.LoginId != "" {
		zap.L().Error("用户已存在", zap.Error(err))
		myCode = code.CheckUserIsExist
		IsExist = false
		return IsExist, myCode, errMsg.UserIsExist
	} else if user.Tel != 0 {
		zap.L().Error("手机号已注册", zap.Error(err))
		myCode = code.TelIsUsed
		IsExist = false
		return IsExist, myCode, errMsg.TelIsUsed
	}
	IsExist = true
	return IsExist, code.SUCCESS, nil
}

// RegisterUserInfo 注册用户信息
func RegisterUserInfo(p *models.User) (code.MyCode, error) {
	p.Password = MD5(p.Password)
	var user = models.User{
		UUID:     p.UUID,
		Name:     p.Name,
		Tel:      p.Tel,
		LoginId:  p.LoginId,
		Password: p.Password,
	}
	err := DB.Debug().Select(
		"uuid", "name", "tel", "login_id", "password",
	).
		Create(&user).
		Error
	if err != nil {
		zap.L().Error("创建用户失败", zap.Error(err))
		return code.InstallUserInfoFailed, err
	}
	return code.SUCCESS, nil
}

// CheckUserInfo 检查用户信息--> 登录验证
func CheckUserInfo(p *models.User) (err error) {
	var user models.User
	if err = DB.Select("password", "uuid").
		Where("login_id = ?", p.LoginId).
		Find(&user).
		Error; err != nil {
		zap.L().Error("检查用户信息失败", zap.Error(err))
		return errMsg.CheckFailed
	}
	if user.Password != MD5(p.Password) {
		zap.L().Error("用户名账号密码不正确")
		return errMsg.LoginIdOrPasswordIsFailed
	}
	return
}

// CheckUser 查询用户
func CheckUser(p *models.User) (*models.CheckUser, code.MyCode, error) {
	// 先定义一个user用来存储信息
	var user models.User
	IsExist, myCode, err := CheckUserIsExist(p.LoginId, p.Tel)
	if IsExist {
		myCode = code.UserIsNotExist
		err = errMsg.UserIsNotExist
		return nil, myCode, err
	}
	// 去数据库中查询是否存在这个用户,如果存在继续下一步的操作
	// 如果用户存在，那么返回这个用户的基本信息，名字，签名，头像等信息
	if p.Tel != 0 {
		if err := DB.Debug().Select("name", "login_id", "sing", "photo", "tel", "status").
			Where("tel = ?", p.Tel).Or("login_id = ?", p.LoginId).
			Find(&user).
			Error; err != nil {
			fmt.Println(err)
			return nil, code.UserIsNotExist, err
		}
	} else {
		if err := DB.Debug().
			Select("name", "sing", "login_id", "photo", "tel", "account", "status").
			Where("tel = ?", p.Tel).Or("login_id = ?", p.LoginId).
			Find(&user).
			Error; err != nil {
			fmt.Println(err)
			return nil, code.UserIsNotExist, err
		}
	}
	return &models.CheckUser{
		Name:    user.Name,
		Sing:    user.Sing,
		Photo:   user.Sing,
		Address: user.Address,
		Tel:     user.Tel,
		Status:  user.Status,
		LoginId: p.LoginId,
	}, code.SUCCESS, nil
}

// MD5 密码验证
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
