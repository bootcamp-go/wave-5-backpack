package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
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
			ctx.JSON(404, web.NewResponse(401, nil, "no tiene permisos para realizar la peticion solicitada"))
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "el id ingresado es invalido"))
			return
		}
		var req reqPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(404, web.NewResponse(400, nil, "El nombre es un campo requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido es un campo requerido"))
		}

		user, error := c.service.UpdateNameAndLastName(id, req.Nombre, req.Apellido)
		if error != nil {
			ctx.JSON(404, web.NewResponse(400, nil, error.Error()))
		}
		ctx.JSON(200, user)

	}
}

func (c *Usuarios) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(401, nil, "no tiene permisos para realizar la peticion solicitada"))
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "El id es invalido"))
			return
		}
		err := c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminaod", id), ""))
	}
}

func (c *Usuarios) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(401, nil, "no tiene permisos para realizar la peticion solicitada"))
			return
		}
		id, error := strconv.Atoi(ctx.Param("id"))
		if error != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "el id es invalido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre es un campo requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido es un campo requerido"))
			return
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El email es un campo requerido"))
			return
		}
		if req.Edad < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad es un campo requerido"))
			return
		}
		if req.Altura < 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La altura es un campo requerido"))
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha es un campo requerido"))
			return
		}

		user, err := c.service.Update(id, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, ""))

	}
}

func (c *Usuarios) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(401, nil, "no tiene permisos para realizar la peticion solicitada"))
			return
		}

		u, erro := c.service.GetAll()
		if erro != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "error al obtener los datos"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *Usuarios) Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(400, nil, "no tiene permisos para realizar la peticion solicitada"))
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
				ctx.JSON(404, web.NewResponse(404, nil, result))
			}
		}
		user, err := c.service.Guardar(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, ""))

	}

}
