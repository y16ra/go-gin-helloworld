package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Check() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info().Msg("check start.")
		ctx.Next()
		log.Info().Msg("check end.")

	}
}
