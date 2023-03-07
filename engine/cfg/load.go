package cfg

import (
	"gopkg.in/ini.v1"
)

func LoadIni(filename string) (*ini.File, error) {
	return ini.Load(filename)
}
