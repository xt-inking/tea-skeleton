package v1

import (
	"context"

	"github.com/tea-frame-go/tea/teaerrors"
	"github.com/tea-frame-go/tea/tealog"
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver/teahandler"
	"github.com/tea-frame-go/tea/teatypes"
	"github.com/tea-frame-go/tea/teavalid"
)

func Register(t *teaserver.RouterTree, passport Passport) {
	logger := tealog.New(
		tealog.NewRecordHandlerText(),
		tealog.NewWriterCloserFile(tealog.FileDir, "http-v1-passport"),
	)
	errorHandler := func(r *teaserver.Request, err teaerrors.Error) {
		logger.Error(r.Context(), err.ErrorStack())
	}
	t.Tree("/passport", func(t *teaserver.RouterTree) {
		// 不鉴权
		t.Tree("/", func(t *teaserver.RouterTree) {
			// 注册
			t.Post("/register", teahandler.Handler(passport.Register, errorHandler))
		})
		// 鉴权
		t.Tree("/", func(t *teaserver.RouterTree) {
			t.Middleware()
		})
		// 半鉴权
		t.Tree("/", func(t *teaserver.RouterTree) {
			t.Middleware()
		})
	})
}

type Passport interface {
	Register(ctx context.Context, req *RegisterReq) (res teatypes.Result[RegisterRes], err teaerrors.Error)
}

type RegisterReq struct{}

func (req *RegisterReq) Validate() (err teavalid.Error) {
	return
}

type RegisterRes struct{}
