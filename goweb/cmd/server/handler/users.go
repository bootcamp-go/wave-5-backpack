package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Edad           int     `json:"edad"`
	Nombre         string  `json:"nombre"`
	Apellido       string  `json:"apellido"`
	Email          string  `json:"email"`
	Fecha_creacion string  `json:"fecha_creacion"`
	Altura         float64 `json:"altura"`
	Activo         bool    `json:"activo"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Credencials"))
			return
		}

		us, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Credencials"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}

		us, err := u.service.Store(req.Edad, req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, us, ""))
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Credencials"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Id"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Edad required"))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Nombre required"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Apellido required"))
			return
		}
		if req.Email == "" {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Email required"))
			return
		}
		if req.Fecha_creacion == "" {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Fecha_creacion required"))
			return
		}
		if req.Altura == 0 {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Altura required"))
			return
		}
		u, err := u.service.Update(id, req.Edad, req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Credencials"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Id"))
			return
		}
		err = u.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("User with id %v deleted", id), ""))
	}
}

func (u *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, "", "Error: Invalid Credencials"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Invalid Id"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Edad required"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(404, web.NewResponse(404, "", "Error: Field Apellido required"))
			return
		}
		u, err := u.service.UpdateLastNameAndAge(id, req.Edad, req.Apellido)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, "", err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}
