package middleware

import (
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver"
	"github.com/tea-frame-go/tea/teanet/teahttp/teaserver/teamiddleware"

	"tea-skeleton/app/tea-skeleton/internal/config"
)

func Cors() func(next teaserver.HandlerFunc) teaserver.HandlerFunc {
	return teamiddleware.Cors(
		config.Config.Cors.AllowOrigins,
		config.Config.Cors.AllowHeaders,
	)
}
