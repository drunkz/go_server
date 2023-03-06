package main

import (
	"fmt"

	"dkz.com/engine"
)

type DBServer struct {
	engine.BaseServer
}

func (d *DBServer) OnInit() {
	fmt.Println("OnInit")
}

func (d *DBServer) OnStart() {
	fmt.Println("OnStart")
}

func (d *DBServer) OnClose() {
	fmt.Println("OnClose")
}
