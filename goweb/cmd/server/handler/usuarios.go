package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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

type reqPatch struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" `
	Edad          int     `json:"edad" `
	Altura        float64 `json:"altura" `
	Activo        bool    `json:"activo" `
	FechaCreacion string  `json:"fecha_de_creacion" `
}

type Usuarios struct {
	service usuarios.Service
}

func NewUsuario(u usuarios.Service) *Usuarios {
	return &Usuarios{service: u}
}

func (c *Usuarios) UpdateNameAndLastName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, gin.H{
				"error": "el id es invalido",
			})
			return
		}
		var req reqPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"eror": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"eror": "El nombre es un campo requerido"})
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"eror": "El apellido es un campo requerido"})
		}

		user, error := c.service.UpdateNameAndLastName(id, req.Nombre, req.Apellido)
		if error != nil {
			ctx.JSON(400, gin.H{"eror": error.Error()})
		}
		ctx.JSON(200, user)

	}
}

func (c *Usuarios) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, gin.H{
				"error": "el id es invalido",
			})
			return
		}
		err := c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminaod", id)})
	}
}

func (c *Usuarios) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, gin.H{
				"error": "el id es invalido",
			})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"eror": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"eror": "El nombre es un campo requerido"})
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"eror": "El apellido es un campo requerido"})
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"eror": "El email es un campo requerido"})
		}
		if req.Edad < 0 {
			ctx.JSON(400, gin.H{"eror": "La edad debe ser un nro positivo"})
		}
		if req.Altura < 0.0 {
			ctx.JSON(400, gin.H{"eror": "La altura debe ser un nro positivo"})
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, gin.H{"eror": "La fecha creacion es un campo obligatorio"})
		}

		user, error := c.service.Update(id, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if error != nil {
			ctx.JSON(400, gin.H{"eror": error.Error()})
		}
		ctx.JSON(200, user)

	}
}

func (c *Usuarios) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
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
		if token != os.Getenv("TOKEN") {
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
