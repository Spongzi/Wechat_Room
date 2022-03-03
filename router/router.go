package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat_room/conf"
	"wechat_room/logger"
)

func InitRouter() {
	gin.SetMode(conf.ServerConn.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), gin.Recovery())
	v1 := r.Group("api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "hahah",
			})
		})
	}
	panic(r.Run(conf.ServerConn.HttpPort))
}
