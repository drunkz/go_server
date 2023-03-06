package cfg

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func Load(filename string) error {
	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("[%s]文件加载失败，%s\n", filename, err)
		return err
	}
	G_BaseConfig.ServerName = cfg.Section("Server").Key("Name").String()
	G_BaseConfig.ServerAddr = cfg.Section("Server").Key("Addr").String()
	G_BaseConfig.ServerPort = uint16(cfg.Section("Server").Key("Port").MustUint(7000))
	G_BaseConfig.LogDir = cfg.Section("Log").Key("Dir").String()
	G_BaseConfig.LogLevel = int8(cfg.Section("Log").Key("Level").MustInt(0))
	fmt.Println(G_BaseConfig)
	return nil
}
