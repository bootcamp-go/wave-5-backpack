package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type Usuarios struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}
