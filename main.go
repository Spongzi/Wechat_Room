package main

import (
	"fmt"
	"go.uber.org/zap"
	"wechat_room/conf"
	"wechat_room/dao"
	"wechat_room/logger"
	"wechat_room/middleware/trans"
	"wechat_room/middleware/uuid"
	"wechat_room/router"
)

// 主程序入口
func main() {
	// 初始配置文件
	if err := conf.InitConf(); err != nil {
		zap.L().Error("Conf Init Failed！", zap.Error(err))
		return
	}
	fmt.Println("测试数据", conf.LogConn.FileName)
	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		zap.L().Error("Logger Init Failed！", zap.Error(err))
		return
	}
	// 初始化雪花算法
	if err := uuid.Init(conf.ServerConn.StartTime, 1); err != nil {
		zap.L().Error("Mysql Init Failed！", zap.Error(err))
		return
	}
	// 初始化数据库连接
	if err := dao.InitMysql(); err != nil {
		zap.L().Error("Mysql Init Failed！", zap.Error(err))
		return
	}
	// 初始化翻译器
	if err := trans.InitTrans("zh"); err != nil {
		zap.L().Error("Mysql Init Failed！", zap.Error(err))
		return
	}
	// 初始化路由
	router.InitRouter()
}
