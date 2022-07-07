package main

import (
	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int     `json:"-"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo,omitempty"`
	Cantidad int     `json:"cantidad" binding:"required"`
	Precio   float64 `json:"precio" binding:"required"`
}

var products []request
var lastID int

func main() {
	server := gin.Default()

	pr := server.Group("/productos")
	pr.POST("/", GuardarWithAuthorize())

	server.Run(":8080")
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++

		req.ID = lastID
		products = append(products, req)

		c.JSON(200, req)
	}
}

func GuardarWithAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		// Validar token
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		// Si el token fue valido, avanzo
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.ID = lastID

		products = append(products, req)

		c.JSON(200, products)
	}

}
