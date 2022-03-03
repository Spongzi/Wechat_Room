package main

import (
	"go.uber.org/zap"
	"wechat_room/conf"
	"wechat_room/logger"
	"wechat_room/mysql"
	"wechat_room/router"
)

// 主程序入口
func main() {
	// 初始配置文件
	if err := conf.InitConf(); err != nil {
		zap.L().Error("Conf Init Failed！", zap.Error(err))
		return
	}
	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		zap.L().Error("Logger Init Failed！", zap.Error(err))
		return
	}
	// 初始化数据库连接
	if err := mysql.InitMysql(); err != nil {
		zap.L().Error("Mysql Init Failed！", zap.Error(err))
		return
	}
	// 初始化路由
	router.InitRouter()
}
