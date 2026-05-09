package http

import (
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"

	"tea-skeleton/app/tea-skeleton/internal/config"
)

func Serve() {
	s := teaserver.New(&config.Config.Server)
	s.Router().Tree("/", func(t *teaserver.RouterTree) {
	})
	s.Serve()
}
