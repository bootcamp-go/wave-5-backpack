package controlador

import (
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Edad     int    `json:"edad"`
	Altura   int    `json:"altura"`
	Activo   bool   `json:"activo"`
}
type Usuario struct {
	service usuarios.Servicio
}

func NewControlador(us usuarios.Servicio) *Usuario {
	return &Usuario{
		service: us,
	}
}
func (u *Usuario) VerUsuarios() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		usuarios, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuarios)
	}
}
func (u *Usuario) AgregarUsuarios() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuarios, err := u.service.Save(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuarios)
	}
}
