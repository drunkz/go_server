package engine

import (
	"log"

	"dkz.com/engine/cfg"
	slog "dkz.com/engine/log"
	"go.uber.org/zap"
)

const CONFIG_FILE_NAME = "config.ini"

type IBaseServer interface {
	OnInit()
	OnStart()
	OnClose()
}

type TBaseServer struct {
	iBaseServer IBaseServer
	BaseConfig  cfg.BaseConfig
	Slog        *zap.Logger
}

func (b *TBaseServer) InitServer(iBaseServer IBaseServer) {
	b.iBaseServer = iBaseServer
	// 加载基本配置
	cfg, err := cfg.LoadIni(CONFIG_FILE_NAME)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	b.BaseConfig.Cfg = cfg
	b.BaseConfig.ServerName = cfg.Section("Server").Key("Name").String()
	b.BaseConfig.ServerAddr = cfg.Section("Server").Key("Addr").String()
	b.BaseConfig.ServerPort = uint16(cfg.Section("Server").Key("Port").MustUint(7000))
	b.BaseConfig.LogDir = cfg.Section("Log").Key("Dir").String()
	b.BaseConfig.LogLevel = int8(cfg.Section("Log").Key("Level").MustInt(0))
	// 初始化日志
	b.Slog = slog.InitLog(b.BaseConfig)
	b.iBaseServer.OnInit()
	b.iBaseServer.OnStart()
}
