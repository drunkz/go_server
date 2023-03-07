package engine

import (
	"fmt"
	"log"

	"dkz.com/engine/cfg"
	g "dkz.com/engine/global"
	slog "dkz.com/engine/log"
	snet "dkz.com/engine/net"
)

const CONFIG_FILE_NAME = "config.ini"

type IBaseServer interface {
	OnInit()  // 服务未启动初始化时
	OnStart() // 服务启动后运行时
	OnClose() // 服务即将关闭时
}

type TBaseServer struct {
	iBaseServer IBaseServer
	BaseConfig  cfg.BaseConfig
}

func (b *TBaseServer) InitServer(iBaseServer IBaseServer) {
	b.iBaseServer = iBaseServer
	// 加载基本配置
	cfg, err := cfg.LoadIni(CONFIG_FILE_NAME)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	b.BaseConfig.Cfg = cfg
	b.BaseConfig.ServerName = cfg.Section("Server").Key("Name").MustString("Server")
	b.BaseConfig.ServerAddr = cfg.Section("Server").Key("Addr").MustString("127.0.0.1")
	b.BaseConfig.ServerPort = uint16(cfg.Section("Server").Key("Port").MustUint(7000))
	b.BaseConfig.LogFileName = cfg.Section("Log").Key("FileName").MustString("log/log.log")
	b.BaseConfig.LogLevel = int8(cfg.Section("Log").Key("Level").InInt(-1, []int{0, 1, 2, 3, 4, 5}))
	b.BaseConfig.LogMaxSize = uint8(cfg.Section("Log").Key("MaxSize").MustUint(5))
	b.BaseConfig.LogMaxAge = uint8(cfg.Section("Log").Key("MaxAge").MustUint(28))
	b.BaseConfig.LogMaxBackups = uint8(cfg.Section("Log").Key("MaxBackups").MustUint(10))
	b.BaseConfig.LogCompress = cfg.Section("Server").Key("Compress").MustBool(false)
	// 初始化日志
	g.Log = slog.InitLog(b.BaseConfig)
	// 初始化网络
	listener, err := snet.Listen(&b.BaseConfig)
	if err != nil {
		g.Log.Fatal(err.Error())
	}
	fmt.Println(b.BaseConfig.ServerPort)
	b.iBaseServer.OnInit()
	listener.Accept()
	b.iBaseServer.OnStart()
}
