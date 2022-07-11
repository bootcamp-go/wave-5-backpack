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
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, web.ERR_TOKEN_INVALID))
			return
		}
		usuarios, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(500, nil, web.ERR_BAD_INTERNAL))
			return
		}
		ctx.JSON(200, usuarios)
	}
}

func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, web.ERR_TOKEN_INVALID))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, web.ERR_BAD_REQUEST))
			return
		}
		user, err := c.service.Store(req.Age, req.Name, req.LastName, req.Email, req.Estatura)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, web.ERR_BAD_INTERNAL))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.ERR_TOKEN_INVALID))
			return
		}
		var req domain.Usuario
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, web.ERR_BAD_REQUEST))
			return
		}
		user, err := c.service.Update(id, req.Age, req.Names, req.LastName, req.Email, req.DateCreated, req.Estatura, req.IsActivo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, web.ERR_BAD_INTERNAL))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) PatchLastNameAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, web.ERR_TOKEN_INVALID))
			return
		}
		var req LastNameAgePatchRequest
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, web.ERR_BAD_REQUEST))
			return
		}
		user, err := c.service.UpdateLastNameAndAge(id, req.Age, req.LastName)
		if err != nil {
			ctx.JSON(404, web.NewResponse(500, nil, web.ERR_BAD_INTERNAL))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, web.ERR_TOKEN_INVALID))
			return
		}
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest,
				web.NewResponse(401, nil, web.ERR_ID_INVALID))
			return
		}
		errDelete := c.service.Delete(id)
		if errDelete != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, web.ERR_BAD_INTERNAL))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, "Se eliminó el usuario correctamente.", ""))
	}
}
