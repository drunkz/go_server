package config

import "gopkg.in/ini.v1"

type BaseConfig struct {
	ServerName    string
	ServerAddr    string
	ServerPort    uint16
	LogFileName   string // 日志文件
	LogLevel      int8   // 日志等级
	LogMaxSize    uint8  // 单个日志文件允许大小
	LogMaxAge     uint8  // 最大保留天数
	LogMaxBackups uint8  // 最大保留数量
	LogCompress   bool   // 是否压缩保留
	Ini           *ini.File
}

const BASE_CONFIG_FILE = "config.ini"
