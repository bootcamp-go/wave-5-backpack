package main

import (
	"goweb/productos/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	internal.SaludoGet(router)
	internal.ProductosGet(router)
	internal.ProductosPost(router)

	router.Run()
}
