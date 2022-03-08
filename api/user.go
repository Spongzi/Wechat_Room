package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"wechat_room/logic"
	"wechat_room/middleware/trans"
	"wechat_room/models"
	"wechat_room/utile/code"
	"wechat_room/utile/resp"
)

// Register 注册
func Register(c *gin.Context) {
	// 绑定输入的信息
	p := new(models.PramsRegister)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("绑定数据失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, code.InvalidParamFailed)
		}
		resp.ResponseErrorWithMsg(c, code.InvalidParamFailed, errs.Translate(trans.Trans))
		return
	}
	// 逻辑运算处理
	myCode, err := logic.Register(p)
	if err != nil {
		resp.ResponseError(c, myCode)
		return
	}
	resp.ResponseSuccess(c, "恭喜注册成功")
}

// Login 登录
func Login(c *gin.Context) {
	// 绑定输入的信息
	p := new(models.PramsLoginUser)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("绑定数据失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, code.InvalidParamFailed)
		}
		resp.ResponseErrorWithMsg(c, code.InvalidParamFailed, errs.Translate(trans.Trans))
		return
	}
	// 逻辑执行, 并且获取token
	myCode, token, err := logic.Login(p)
	if err != nil {
		resp.ResponseError(c, myCode)
		return
	}
	// 返回结果
	resp.ResponseSuccess(c, token)
}

// CheckUser 查询好友
func CheckUser(c *gin.Context) {
	// 首先，绑定要查询人的信息，可以是用户账号，也可以是用户的手机号
	p := new(models.PramsCheckUser)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("绑定信息失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, code.InvalidParamFailed)
		}
		resp.ResponseErrorWithMsg(c, code.InvalidParamFailed, errs.Translate(trans.Trans))
		return
	}
	// 逻辑判断
	user, myCode, err := logic.CheckUser(p)
	if err != nil {
		resp.ResponseError(c, myCode)
		return
	}
	resp.ResponseSuccess(c, user)
}

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	var friend models.PramsAddFriend
	// 首先绑定输入的用户
	if err := c.ShouldBindJSON(&friend); err != nil {
		zap.L().Error("绑定数据失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			resp.ResponseError(c, code.InvalidParamFailed)
			return
		}
		resp.ResponseErrorWithMsg(c, code.InvalidParamFailed, errs.Translate(trans.Trans))
		return
	}
	// 逻辑操作, 发送添加数据
	logic.AddFriend()
	// 返回结果
}
