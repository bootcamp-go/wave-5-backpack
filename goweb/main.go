package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Altura   float64 `json:"altura" binding:"required"`
	Activo   bool    `json:"activo" binding:"required"`
	Fecha    string  `json:"fecha" binding:"required"`
}

var Usuarios []Usuario

var lastID int = 3
var tokenAcceso string = "123456"

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "Hola Diego")
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": Usuarios,
	})
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Usuario
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++

		req.Id = lastID
		Usuarios = append(Usuarios, req)

		c.JSON(200, req)
	}
}

func GuardarConToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Usuario
		// Validar token
		token := c.GetHeader("token")
		if token != tokenAcceso {
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
		req.Id = lastID

		Usuarios = append(Usuarios, req)

		c.JSON(200, req)
	}

}

func main() {

	Usuarios = append(Usuarios, Usuario{1, "Diego", "Palacios", "d@example.com", 1.68, true, "21/06/2022"})
	Usuarios = append(Usuarios, Usuario{2, "Fernndo", "Palacios", "f@example.com", 1.70, true, "21/06/2022"})
	Usuarios = append(Usuarios, Usuario{3, "Cesar", "Parrado", "c@example.com", 1.80, true, "21/06/2022"})

	server := gin.Default()
	server.GET("", PaginaPrincipal)
	pr := server.Group("/usuarios")
	pr.GET("/", GetAll)
	pr.POST("/", GuardarConToken())
	pr.POST("/sintoken", Guardar())

	server.Run()

}
