package code

type MyCode int

const (
	SUCCESS                   MyCode = 1000
	ERROR                            = 5000
	CodeServerBusy                   = 3000
	TOKENAUTOISNIL                   = 2003
	TOKENAUTOISFAILED                = 2004
	TOKENAUTOISNOTEXIT               = 2005
	CheckUserIsExist                 = 3001
	LoginIdOrPasswordIsFailed        = 3002
	InstallUserInfoFailed            = 3003
	InvalidParamFailed               = 3004
	UserIsNotExist                   = 3005
	CheckFailed                      = 3006
	TelIsUsed                        = 3007
)

var codeMsg = map[MyCode]string{
	SUCCESS:                   "请求成功",
	ERROR:                     "请求失败",
	CodeServerBusy:            "服务器繁忙",
	TOKENAUTOISNIL:            "请求头中auth为空",
	TOKENAUTOISFAILED:         "请求头中auth格式有误",
	TOKENAUTOISNOTEXIT:        "无效的Token",
	CheckUserIsExist:          "用户已存在",
	LoginIdOrPasswordIsFailed: "用户名或密码错误",
	InstallUserInfoFailed:     "创建用户失败",
	InvalidParamFailed:        "绑定数据失败",
	UserIsNotExist:            "用户不存在",
	CheckFailed:               "检查用户是否存在错误",
	TelIsUsed:                 "手机号已经被注册",
}

func (m MyCode) GetMsg() string {
	msg, ok := codeMsg[m]
	if !ok {
		msg = codeMsg[CodeServerBusy]
	}
	return msg
}
