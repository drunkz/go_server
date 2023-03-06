package main

import (
	"fmt"

	"dkz.com/engine"
)

var G_DBServer *TDBServer

type TDBServer struct {
	engine.TBaseServer
}

func (d *TDBServer) OnInit() {
	fmt.Println("OnInit")
}

func (d *TDBServer) OnStart() {
	fmt.Println("OnStart")
}

func (d *TDBServer) OnClose() {
	fmt.Println("OnClose")
}
