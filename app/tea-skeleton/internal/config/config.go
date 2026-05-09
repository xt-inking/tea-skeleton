package config

import (
	"github.com/tea-frame-go/tea/teaconfig"
)

func Init() {
	teaconfig.Decode(&Config)
}

var Config = config{
	Server: teaconfig.NewServerConfig(),
	Cors:   teaconfig.NewCorsConfig(),
}

type config struct {
	Server teaconfig.ServerConfig
	Cors   teaconfig.CorsConfig
}
