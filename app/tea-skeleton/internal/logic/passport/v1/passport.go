package v1

import (
	"context"

	"github.com/tea-frame-go/tea/teaerrors"
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"
	"github.com/tea-frame-go/tea/teatypes"

	v1 "tea-skeleton/app/tea-skeleton/api/passport/v1"
)

func Register(t *teaserver.RouterTree) {
	v1.Register(t, passport{})
}

type passport struct{}

func (passport) Register(ctx context.Context, req *v1.RegisterReq) (res teatypes.Result[v1.RegisterRes], err teaerrors.Error) {
	return
}
