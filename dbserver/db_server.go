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

func (d *TDBServer) OnInit() {
	g.Log.Debug("ttttttttttttt")
	g.Log.Error("asdfsdf看看中文显示效果怎么样adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显示效果怎adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显示效么样adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显怎么样adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显怎么样adsf3545341ad")
	g.Log.Info("asdfsdf看看中文显示效果怎么样adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显示效果怎么样adsf3545341ad")
	g.Log.Debug("asdfsdf看看中文显示效果怎么样adsf3545341ad")
}

func (d *TDBServer) OnStart() {
	fmt.Println("OnStart")
}

func (d *TDBServer) OnClose() {
	defer g.Log.Sync()
	fmt.Println("OnClose")
}
