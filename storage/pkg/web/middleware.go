package web

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontró el token")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewResponse(http.StatusUnauthorized, nil, "no ingresó el token"))
			return
		}

		if requiredToken != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		ctx.Next()
	}
}
