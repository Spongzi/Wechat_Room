package errMsg

import "errors"

var (
	CheckFailed               = errors.New("查询失败")
	UserIsNotExist            = errors.New("用户不存在")
	LoginIdOrPasswordIsFailed = errors.New("用户名或密码错误")
)
