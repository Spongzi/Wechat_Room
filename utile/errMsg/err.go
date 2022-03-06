package errMsg

import "errors"

var (
	CheckFailed               = errors.New("查询失败")
	UserIsExist               = errors.New("用户已经存在")
	LoginIdOrPasswordIsFailed = errors.New("用户名或密码错误")
	TelIsUsed                 = errors.New("手机号已被注册")
	UserIsNotExist            = errors.New("用户不存在")
)
