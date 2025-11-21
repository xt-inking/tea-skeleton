package passport

import (
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"

	v1 "tea-skeleton/app/tea-skeleton/internal/logic/passport/v1"
)

func RegisterV1(t *teaserver.RouterTree) {
	v1.Register(t)
}
