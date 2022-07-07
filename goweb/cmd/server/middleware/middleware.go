package middleware

import "github.com/gin-gonic/gin"

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token != "Bearer 123456" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}
}
