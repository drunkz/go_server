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

	"dkz.com/engine/config"
	g "dkz.com/engine/global"
	slog "dkz.com/engine/log"
	"dkz.com/engine/mode"
	snet "dkz.com/engine/net"
	"dkz.com/engine/system"
)

type IBaseServer interface {
	OnStart()
	OnStop()
}

type TBaseServer struct {
	iBaseServer IBaseServer
	BaseConfig  config.BaseConfig
	Server      net.Listener
}

func (b *TBaseServer) InitServer(iBaseServer IBaseServer) {
	g.IsDebugMode = mode.IsDebugMode()
	b.iBaseServer = iBaseServer
	// 加载基本配置
	err := config.InitBaseConfig(&b.BaseConfig)
	if err != nil {
		log.Fatalln(err)
	}
	// 初始化日志
	g.Log = slog.InitLog(b.BaseConfig)
	// Windows 平台设置
	if !g.IsDebugMode && runtime.GOOS == "windows" {
		if err = system.InitPlatform(&b.BaseConfig); err != nil {
			g.Log.Fatal(err.Error())
		}
	}
	// 初始化网络
	b.Server, err = snet.Listen(&b.BaseConfig)
	if err != nil {
		g.Log.Fatal(err.Error())
	}
	fmt.Printf("[%s]开始运行...按下Ctrl+C终止...\n", b.BaseConfig.ServerName)
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
