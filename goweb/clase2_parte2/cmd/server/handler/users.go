package handler

import (
	"github.com/gin-gonic/gin"
	"goweb/clase2_parte2/internal/users"
)

// Se debe generar la estructura request
type request struct {
	Nombre 			string	`json:"nombre"`
	Apellido 		string	`json:"apellido"`
	Email 			string	`json:"email"`
	Edad 			int		`json:"edad"`
	Altura 			float64	`json:"altura"`
	Activo			bool	`json:"activo"`
	FechaCreacion 	string	`json:"fecha_creacion"`
}

// Se debe generar la estructura del controlador que tenga como campo el servicio
type User struct {
	service users.Service
}

// Se debe generar la función que retorne el controlador
func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

// Se deben generar todos los métodos correspondientes a los endpoints
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "El token no es válido"})
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{ "error": "El token no es válido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}