package config

import (
	"github.com/tea-frame-go/tea/teaconfig"
)

func Init() {
	teaconfig.Decode(&Config)
}

var Config = config{
	Cors: cors{
		AllowOrigins: []string{"localhost:8080"},
		AllowHeaders: "",
	},
	Server: &teaconfig.Server,
}

type config struct {
	Cors   cors
	Server *teaconfig.ServerConfig
}

type cors struct {
	AllowOrigins []string
	AllowHeaders string
}
