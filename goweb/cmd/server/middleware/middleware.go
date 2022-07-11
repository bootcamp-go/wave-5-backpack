package middleware

import (
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
		return
	}
}
