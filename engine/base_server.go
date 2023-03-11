package engine

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"dkz.com/engine/cfg"
	g "dkz.com/engine/global"
	slog "dkz.com/engine/log"
	"dkz.com/engine/mode"
	snet "dkz.com/engine/net"
	"dkz.com/engine/system"
)

const CONFIG_FILE_NAME = "config.ini"

type IBaseServer interface {
	OnStart()
	OnStop()
}

type TBaseServer struct {
	iBaseServer IBaseServer
	BaseConfig  cfg.BaseConfig
	Server      net.Listener
}

func (b *TBaseServer) InitServer(iBaseServer IBaseServer) {
	g.IsDebugMode = mode.IsDebugMode()
	b.iBaseServer = iBaseServer
	// 加载基本配置
	cfg, err := cfg.LoadIni(CONFIG_FILE_NAME)
	if err != nil {
		log.Fatalln(err)
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
	// Windows 平台设置
	if !g.IsDebugMode && runtime.GOOS == "windows" {
		if _, err = system.NewModule(system.Kernel32Dll, "GetConsoleWindow", "SetConsoleTitleW"); err != nil {
			g.Log.Fatal(err.Error())
		}
		// 设置控制台窗口标题
		if err = system.SetTitle(b.BaseConfig.ServerName); err != nil {
			g.Log.Fatal(err.Error())
		}
		if err = system.InitConsoleHandle(); err != nil {
			g.Log.Fatal(err.Error())
		}
		// 禁止快速编辑模式
		if err = system.DisableQuickEdit(); err != nil {
			g.Log.Fatal(err.Error())
		}
		if _, err = system.NewModule(system.User32Dll, "GetSystemMenu", "RemoveMenu"); err != nil {
			g.Log.Fatal(err.Error())
		}
		// 移除关闭和最大化按钮
		if err = system.RemoveMenu(); err != nil {
			g.Log.Fatal(err.Error())
		}
		system.FreeModule()
	}
	// 初始化网络
	b.Server, err = snet.Listen(&b.BaseConfig)
	if err != nil {
		g.Log.Fatal(err.Error())
	}
	b.iBaseServer.OnStart()
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go b.serverAccept(ctx)
	<-ctx.Done()
	b.iBaseServer.OnStop()
	b.Server.Close()
	g.Log.Sync()
}

func (b *TBaseServer) serverAccept(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := b.Server.Accept()
			if err != nil {
				if !strings.Contains(err.Error(), "use of closed network connection") {
					g.Log.Error(err.Error())
				}
				continue
			}
			go b.processClient(ctx, conn)
		}
	}
}

func (b *TBaseServer) processClient(ctx context.Context, conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			len, err := conn.Read(buffer)
			if err != nil && err != io.EOF {
				g.Log.Error(err.Error())
				continue
			}
			if len == 0 {
				continue
			}
			fmt.Printf("收到：%s\t%s\n", conn.RemoteAddr().String(), string(buffer[:len]))
			conn.Write(buffer[:len])
			//conn.Close()
		}
	}
}
