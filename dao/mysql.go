package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wechat_room/conf"
	"wechat_room/models"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlConn.User,
		conf.MysqlConn.Password,
		conf.MysqlConn.Url,
		conf.MysqlConn.Port,
		conf.MysqlConn.DBName,
	)
	if DB, err = gorm.Open(mysql.Open(dsn)); err != nil {
		zap.L().Error("MySQL连接失败", zap.Error(err))
		return err
	}
	if err = DB.AutoMigrate(&models.User{}); err != nil {
		zap.L().Error("数据库迁移失败", zap.Error(err))
		return err
	}
	return
}
