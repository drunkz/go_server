package main

import (
	"fmt"

	"dkz.com/engine"
	g "dkz.com/engine/global"
)

var G_DBServer *TDBServer

type TDBServer struct {
	engine.TBaseServer
}

func (d *TDBServer) OnStart() {
	g.Log.Debug(fmt.Sprintf("绑定地址：%s:%d", d.BaseConfig.ServerAddr, d.BaseConfig.ServerPort))
}

func (d *TDBServer) OnStop() {
	g.Log.Debug("OnStop")
}
