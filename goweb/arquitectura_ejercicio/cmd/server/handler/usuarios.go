package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/usuarios"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name     string  `json:"nombre"`
	LastName string  `json:"apellido"`
	Email    string  `json:"email"`
	Age      int     `json:"edad"`
	Estatura float64 `json:"altura"`
}
type LastNameAgePatchRequest struct {
	LastName string `json:"apellido"`
	Age      int    `json:"edad"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(u usuarios.Service) *Usuario {
	return &Usuario{
		service: u,
	}
}

func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		usuarios, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Existió un error en la consulta"))
			return
		}
		ctx.JSON(200, usuarios)
	}
}

func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Existió un error en la consulta"))
			return
		}
		user, err := c.service.Store(req.Age, req.Name, req.LastName, req.Email, req.Estatura)
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, "Existió un error en..."))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req domain.Usuario
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Existió un error en la request."))
			return
		}
		user, err := c.service.Update(id, req.Age, req.Names, req.LastName, req.Email, req.DateCreated, req.Estatura, req.IsActivo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, "Existió un problema en la consulta."))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) PatchLastNameAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req LastNameAgePatchRequest
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Existió un error en la request."))
			return
		}
		user, err := c.service.UpdateLastNameAndAge(id, req.Age, req.LastName)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Existió un error en la consulta."))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest,
				web.NewResponse(401, nil, "El ID es inválido"))
			return
		}
		errDelete := c.service.Delete(id)
		if errDelete != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Exisitó un problema en la request."))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, "Se eliminó el usuario correctamente.", ""))
	}
}
