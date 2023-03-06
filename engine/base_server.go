package engine

import (
	"fmt"

	"dkz.com/engine/cfg"
)

type IBaseServer interface {
	OnInit()
	OnStart()
	OnClose()
}

type BaseServer struct {
	IBaseServer
}

func (b *BaseServer) ServerInit() {
	fmt.Println("server_init...")
	err := cfg.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	b.OnStart()
}
