package net

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"dkz.com/engine/cfg"
)

func Listen(config *cfg.BaseConfig) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.ServerAddr, config.ServerPort))
	if err != nil {
		return nil, err
	}
	// 使用 0 号端口时保存实际端口
	if config.ServerPort == 0 {
		sUsePort := listener.Addr().String()
		sUsePort = sUsePort[strings.LastIndex(sUsePort, ":")+1:]
		nUsePort, err := strconv.Atoi(sUsePort)
		if err != nil {
			return nil, err
		}
		config.ServerPort = uint16(nUsePort)
	}
	return listener, nil
}
