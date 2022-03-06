package router

import (
	"github.com/gin-gonic/gin"
	"wechat_room/api"
	"wechat_room/conf"
	"wechat_room/logger"
	"wechat_room/middleware/jwtToken"
)

func InitRouter() {
	gin.SetMode(conf.ServerConn.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), gin.Recovery())
	v1 := r.Group("api/v1")
	{
		v1.POST("/login", api.Login)
		v1.POST("/register", api.Register)
	}
	v1.Use(jwtToken.JWTAuthMiddleware())
	{
		v1.POST("/check", api.CheckUser)
	}
	panic(r.Run(conf.ServerConn.HttpPort))
}
