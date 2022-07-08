package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token != fmt.Sprintf("Bearer %s", os.Getenv("TOKEN")) {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}
}
