package config

import (
	"gopkg.in/ini.v1"
)

func LoadIni(filename string) (*ini.File, error) {
	return ini.Load(filename)
}

func InitBaseConfig(cfg *BaseConfig) error {
	ini, err := LoadIni(BASE_CONFIG_FILE)
	if err != nil {
		return err
	}
	cfg.Ini = ini
	cfg.ServerName = ini.Section("Server").Key("Name").MustString("Server")
	cfg.ServerAddr = ini.Section("Server").Key("Addr").MustString("127.0.0.1")
	cfg.ServerPort = uint16(ini.Section("Server").Key("Port").MustUint(7000))
	cfg.LogFileName = ini.Section("Log").Key("FileName").MustString("log/log.log")
	cfg.LogLevel = int8(ini.Section("Log").Key("Level").InInt(-1, []int{0, 1, 2, 3, 4, 5}))
	cfg.LogMaxSize = uint8(ini.Section("Log").Key("MaxSize").MustUint(5))
	cfg.LogMaxAge = uint8(ini.Section("Log").Key("MaxAge").MustUint(28))
	cfg.LogMaxBackups = uint8(ini.Section("Log").Key("MaxBackups").MustUint(10))
	cfg.LogCompress = ini.Section("Server").Key("Compress").MustBool(false)
	return nil
}
