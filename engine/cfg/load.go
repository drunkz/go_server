package cfg

import (
	"gopkg.in/ini.v1"
)

func LoadIni(filename string) (*ini.File, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
