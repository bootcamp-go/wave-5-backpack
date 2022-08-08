package handler

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

// type requestBD struct {
// 	Nombre        string  `json:"nombre" binding:"required"`
// 	Apellido      string  `json:"apellido" binding:"required"`
// 	Email         string  `json:"email" binding:"required"`
// 	Edad          int     `json:"edad" binding:"required"`
// 	Altura        float64 `json:"altura" binding:"required"`
// 	Activo        bool    `json:"activo" binding:"required"`
// 	FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
// }

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
