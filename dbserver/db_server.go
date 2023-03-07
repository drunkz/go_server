package main

import (
	"dkz.com/engine"
	g "dkz.com/engine/global"
)

var G_DBServer *TDBServer

type TDBServer struct {
	engine.TBaseServer
}

func (d *TDBServer) OnInit() {
	g.Log.Debug("OnInit")
}

func (d *TDBServer) OnStart() {
	g.Log.Debug("OnStart")
}

func (d *TDBServer) OnClose() {
	defer g.Log.Sync()
	g.Log.Debug("OnClose")
}
