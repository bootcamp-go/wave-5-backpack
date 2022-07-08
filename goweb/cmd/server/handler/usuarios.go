package handler

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type request struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int     `json:"edad" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
}

type Usuarios struct {
	service usuarios.Service
}

func NewUsuario(u usuarios.Service) *Usuarios {
	return &Usuarios{service: u}
}

func (c *Usuarios) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345678" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		u, erro := c.service.GetAll()
		if erro != nil {
			ctx.JSON(404, gin.H{
				"error": erro.Error(),
			})
			return
		}
		ctx.JSON(200, u)

	}
}

func (c *Usuarios) Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345678" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("el campo %s es requerido y ", field.Tag())
					} else {
						result += fmt.Sprintf("el campo %s es requerido", field.Tag())
					}
				}
				ctx.JSON(404, result)
			}
		}
		user, error := c.service.Guardar(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if error != nil {
			ctx.JSON(404, gin.H{
				"error": error.Error()})
			return
		}
		ctx.JSON(200, user)

	}

}
