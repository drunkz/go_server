package cfg

import "gopkg.in/ini.v1"

type BaseConfig struct {
	ServerName string
	ServerAddr string
	ServerPort uint16
	LogDir     string // 日志保存目录
	LogLevel   int8   // 日志等级
	Cfg        *ini.File
}
