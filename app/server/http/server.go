package http

import (
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"
)

func Serve() {
	s := teaserver.New()
	s.Router().Tree("/", func(t *teaserver.RouterTree) {
		t.Tree("/", func(t *teaserver.RouterTree) {
		})
	})
	s.Serve()
}
