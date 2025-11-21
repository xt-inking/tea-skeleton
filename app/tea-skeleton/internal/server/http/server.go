package http

import (
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver/teamiddleware"

	"tea-skeleton/app/tea-skeleton/api/passport"
	"tea-skeleton/app/tea-skeleton/internal/server/http/middleware"
)

func Serve() {
	s := teaserver.New()
	s.Router().Tree("/", func(t *teaserver.RouterTree) {
		middlewareRecover := teamiddleware.Recover(s)
		middlewareCors := middleware.Cors()
		t.Tree("/", func(t *teaserver.RouterTree) {
			t.Middleware(
				middlewareRecover,
				middlewareCors,
			)
			t.Tree("/v1", func(t *teaserver.RouterTree) {
				passport.RegisterV1(t)
			})
			t.Tree("/v2", func(t *teaserver.RouterTree) {
			})
		})
		t.Tree("/", func(t *teaserver.RouterTree) {
			t.Middleware(
				middlewareRecover,
			)
			t.Tree("/v1", func(t *teaserver.RouterTree) {
			})
			t.Tree("/v2", func(t *teaserver.RouterTree) {
			})
		})
	})
	s.Serve()
}
