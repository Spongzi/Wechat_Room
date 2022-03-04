package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"wechat_room/utile/code"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  "请求成功",
		"data": data,
	})
}

func ResponseError(c *gin.Context, myCode code.MyCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": myCode,
		"msg":  myCode.GetMsg(),
	})
}

func ResponseErrorWithMsg(c *gin.Context, myCode code.MyCode, msg validator.ValidationErrorsTranslations) {
	c.JSON(http.StatusOK, gin.H{
		"code": myCode,
		"msg":  msg,
	})
}
