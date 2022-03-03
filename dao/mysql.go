package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wechat_room/conf"
	"wechat_room/models"
)

var Db *gorm.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlConn.User,
		conf.MysqlConn.Password,
		conf.MysqlConn.Url,
		conf.MysqlConn.Port,
		conf.MysqlConn.DBName,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("MySQL连接失败", zap.Error(err))
		return err
	}
	// 迁移数据库
	if err = Db.AutoMigrate(&models.User{}, &models.Friend{}); err != nil {
		zap.L().Error("迁移数据库失败", zap.Error(err))
		return
	}
	// 设置连接数
	return
}
