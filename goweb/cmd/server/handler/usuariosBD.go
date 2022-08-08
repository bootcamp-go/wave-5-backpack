package handler

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestBD struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int     `json:"edad" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
}

// type reqPatchBD struct {
// 	Nombre        string  `json:"nombre" binding:"required"`
// 	Apellido      string  `json:"apellido" binding:"required"`
// 	Email         string  `json:"email" `
// 	Edad          int     `json:"edad" `
// 	Altura        float64 `json:"altura" `
// 	Activo        bool    `json:"activo" `
// 	FechaCreacion string  `json:"fecha_de_creacion" `
// }

type UsuariosBD struct {
	serviceBD usuarios.ServiceBD
}

func NewUsuarioBD(u usuarios.ServiceBD) *UsuariosBD {
	return &UsuariosBD{serviceBD: u}
}

func (c *UsuariosBD) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, erro := c.serviceBD.GetByName(ctx.Param("nombre"))
		if erro != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "no existen registros con el nombre indicado"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *UsuariosBD) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestBD
		if err := ctx.ShouldBindJSON(&req); err != nil {
			fmt.Print(err)
		}
		var user domain.Usuarios
		user.Nombre = req.Nombre
		user.Apellido = req.Apellido
		user.Email = req.Email
		user.Edad = req.Edad
		user.Altura = req.Altura
		user.Activo = req.Activo
		user.FechaCreacion = req.FechaCreacion

		user, err := c.serviceBD.StoreBD(user)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, ""))
	}
}
