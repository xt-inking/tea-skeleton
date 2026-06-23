package main

import (
	"tea-skeleton/app/tea-skeleton/internal/config"
	"tea-skeleton/app/tea-skeleton/internal/server/http"
)

func main() {
	config.Init()
	http.Serve()
}
