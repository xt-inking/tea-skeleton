package config

import (
	"github.com/tea-frame-go/tea/teaconfig"
)

func Init() {
	teaconfig.Decode()
}
