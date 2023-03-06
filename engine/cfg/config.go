package cfg

type BaseConfig struct {
	ServerName string
	ServerAddr string
	ServerPort uint16
	LogDir     string // 日志保存目录
	LogLevel   int8   // 日志等级
}

var G_BaseConfig BaseConfig
