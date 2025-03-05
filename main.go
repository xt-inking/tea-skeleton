package main

import (
	"tea-skeleton/app/config"
	"tea-skeleton/app/server/http"
)

func main() {
	config.Init()
	http.Serve()
}
