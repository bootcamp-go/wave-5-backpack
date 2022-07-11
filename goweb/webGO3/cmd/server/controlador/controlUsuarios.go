package controlador

import (
	"fmt"
	"os"
	"strconv"

	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Email          string `json:"email"`
	Fecha_creacion string `json:"fecha_creacion"`
	Edad           int    `json:"edad"`
	Altura         int    `json:"altura"`
	Activo         *bool  `json:"activo"`
}
type Usuario struct {
	service usuarios.Servicio //service = servicio no lo cambie porque ya lo implemente en mas de 3 lugares
}

func NewControlador(us usuarios.Servicio) *Usuario {
	return &Usuario{
		service: us,
	}
}
func (u *Usuario) VerUsuarios() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
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
		if token != os.Getenv("TOKEN") {
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
		usuario, err := u.service.Save(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, *req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuario)
	}
}
func (u *Usuario) ActualizarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := u.service.UpdateUsuario(req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, id, req.Edad, req.Altura, *req.Activo)
		if err != nil {
			ctx.JSON(404, err)
		}
		ctx.JSON(200, usuario)

	}
}

func (u *Usuario) ActualizarAtribUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := u.service.UpdateAtributos(req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, id, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, err)
		}
		ctx.JSON(200, usuario)
	}
}

func (u *Usuario) BorrarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		err = u.service.DeleteUsuario(id)
		if err != nil {
			ctx.JSON(404, err)
		}
		ctx.JSON(200, fmt.Sprint("se borro el usuario de id ", id))
	}
}
