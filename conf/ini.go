package conf

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"os"
)

type Server struct {
	HttpPort  string `ini:"HttpPort"`
	JwtToken  string `ini:"JwtToken"`
	Mode      string `ini:"Mode"`
	StartTime string `ini:"StartTime"`
	MachineID int    `ini:"MachineID"`
}

type Log struct {
	FileName  string `ini:"FileName"`
	MaxSize   int    `ini:"MaxSize"`
	MaxBackUp int    `ini:"MaxBackUp"`
	MaxAge    int    `ini:"MaxAge"`
	Level     string `ini:"Level"`
}

type Mysql struct {
	User        string `ini:"user"`
	Password    string `ini:"password"`
	DBName      string `ini:"dbname"`
	Url         string `ini:"url"`
	Port        string `ini:"port"`
	MaxOpenConn int    `ini:"MaxOpenConn"`
	MaxIdleConn int    `ini:"MaxIdleConn"`
}

// 这里是在设置默认值
var (
	LogConn    = new(Log)
	MysqlConn  = new(Mysql)
	ServerConn = new(Server)
)

func InitConf() (err error) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("Fail to read file", err)
		os.Exit(1)
	}
	err = loadLog(cfg)
	err = loadMysql(cfg)
	err = loadServer(cfg)
	return
}

func loadLog(cfg *ini.File) (err error) {
	// Log 相关的内容
	LogConn = &Log{
		FileName:  "./WeChatLog.log",
		MaxSize:   10,
		MaxAge:    10,
		MaxBackUp: 10,
		Level:     "debug",
	}
	err = cfg.Section("log").MapTo(LogConn)
	if err != nil {
		zap.L().Error("配置错误，请检查", zap.Error(err))
		return
	}
	return
}

func loadMysql(cfg *ini.File) (err error) {
	// Mysql 相关信息
	MysqlConn = &Mysql{
		User:        "root",
		Password:    "123456",
		DBName:      "WeChat",
		Url:         "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 100,
		MaxIdleConn: 60,
	}
	err = cfg.Section("mysql").MapTo(MysqlConn)
	if err != nil {
		zap.L().Error("配置错误，请检查", zap.Error(err))
		return
	}
	return
}

func loadServer(cfg *ini.File) (err error) {
	ServerConn = &Server{
		HttpPort:  ":8080",
		JwtToken:  "WeChatRoom",
		Mode:      "debug",
		StartTime: "2022-03-03",
		MachineID: 1,
	}
	err = cfg.Section("server").MapTo(ServerConn)
	if err != nil {
		zap.L().Error("配置错误，请检查", zap.Error(err))
		return
	}
	return
}
