package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"wechat_room/conf"
)

var DB *sqlx.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlConn.User,
		conf.MysqlConn.Password,
		conf.MysqlConn.Url,
		conf.MysqlConn.Port,
		conf.MysqlConn.DBName,
	)
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("MySQL连接失败", zap.Error(err))
		return err
	}
	// 设置连接数
	DB.SetMaxOpenConns(conf.MysqlConn.MaxOpenConn)
	DB.SetMaxIdleConns(conf.MysqlConn.MaxIdleConn)
	return
}