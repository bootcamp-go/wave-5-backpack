package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InternalError(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{
		"error": "internal server error",
	})
	fmt.Printf("Error: %v\n", err)
}
