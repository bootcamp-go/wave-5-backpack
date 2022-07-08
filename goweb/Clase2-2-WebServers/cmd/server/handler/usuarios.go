package handler

import (
	"fmt"
	"goweb/Clase2-2-WebServers/internal/usuarios"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"binding:"required"`
	Email           string `json:"email"`
	Edad            int    `json:"edad"`
	Altura          int    `json:"altura"`
	Activo          bool   `json:"activo"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUser(p usuarios.Service) *Usuario {
	return &Usuario{
		service: p,
	}
}

func (c *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		user, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, user)
	}
}

func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println("Handler req:")
		fmt.Println(req)
		user, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Handler user:")
		fmt.Println(user)
		ctx.JSON(200, user)
	}
}
