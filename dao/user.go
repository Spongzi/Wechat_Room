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
func CheckUserIsExist(loginId string) (err error) {
	var count int
	sqlStr := "select count(login_Id) from USER where login_id = ?;"
	err = DB.Get(&count, sqlStr, loginId)
	if err != nil {
		zap.L().Error("检查用户是否存在错误", zap.Error(err))
		return errMsg.CheckFailed
	}
	fmt.Println(count)
	if count > 0 {
		zap.L().Error("用户已存在", zap.Error(err))
		return errMsg.UserIsNotExist
	}
	return
}

// RegisterUserInfo 注册用户信息
func RegisterUserInfo(p *models.User) (code.MyCode, error) {
	p.Password = MD5(p.Password)
	sqlStr := "insert into user (uuid, name, tel, login_id, password) values (?, ?, ?, ?, ?);"
	_, err := DB.Exec(sqlStr, p.UUID, p.Name, p.Tel, p.LoginId, p.Password)
	if err != nil {
		zap.L().Error("创建用户失败", zap.Error(err))
		return code.InstallUserInfoFailed, err
	}
	return code.SUCCESS, nil
}

// CheckUserInfo 检查用户信息
func CheckUserInfo(p *models.User) (err error) {
	var password string
	sqlStr := "select password from USER where login_id = ? ;"
	err = DB.Get(&password, sqlStr, p.LoginId)
	if err != nil {
		zap.L().Error("检查用户信息失败", zap.Error(err))
		return errMsg.CheckFailed
	}
	if password != MD5(p.Password) {
		zap.L().Error("用户名账号密码不正确")
		return errMsg.LoginIdOrPasswordIsFailed
	}
	return
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
