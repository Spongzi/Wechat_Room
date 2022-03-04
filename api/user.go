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
