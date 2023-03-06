package main

import (
	"fmt"
	"time"
)

func main() {
	dbServer := DBServer{}
	dbServer.ServerInit()
	time.Sleep(time.Second * 1)
	fmt.Println("aaaaaaaaaaaaaa")
}
