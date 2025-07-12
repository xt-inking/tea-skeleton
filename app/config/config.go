package config

import (
	"github.com/tea-frame-go/tea/teaconfig"
)

var App = app{}

type app struct{}

func Init() {
	teaconfig.Decode(&App)
}
