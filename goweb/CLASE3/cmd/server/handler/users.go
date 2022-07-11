package handler

import (
	"CLASE3/internal/users"
	"CLASE3/pkg/web"
	"os"

	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
}
type Users struct {
	service users.Service
}

func NewUsers(usuario users.Service) *Users {
	return &Users{
		service: usuario,
	}
}

func (c *Users) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.Response(401, nil, "Token Invalido"))
			return
		}
		usuario, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.Response(200, usuario, ""))
	}
}

func (c *Users) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.Response(401, nil, "Token Invalido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		usuario, err := c.service.Store(req.Nombre, req.Apellido, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.Response(200, usuario, ""))
	}
}

func (u *Users) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.Response(401, nil, "Token Invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.Response(400, nil, "El id no es valido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre es requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido es requerido"))
			return
		}
		if req.Edad <= 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad es requerida"))
			return
		}
		if req.Altura <= 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La altura es requerida"))
			return
		}

		u, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))

	}
}

func (u *Users) UpdateApellidoAndEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.Response(401, nil, "Token Invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El id no es valido"))
			return
		}
		type request struct {
			Apellido string `json:"apellido" binding:"required"`
			Edad     int    `json:"edad" binding:"required"`
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido es requerido"))
			return
		}
		if req.Edad <= 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Edad incorrecta"))
			return
		}

		u, err := u.service.UpdateApellidoAndEdad(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))

	}
}

func (u *Users) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(400, web.Response(401, nil, "Token Invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El id no es valido"))
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.Response(200, u, ""))
	}
}
