package handler

import (
	"C4-TM/internal/usuarios"
	"C4-TM/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(s usuarios.Service) *Usuario {
	return &Usuario{service: s}
}

func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		usuarios, err := u.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, "token inválido")) // 500
			return
		}

		if len(usuarios) <= 0 {
			c.JSON(401, web.NewResponse(401, nil, "No se encontraron usuarios."))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, usuarios, ""))
	}
}

func (u *Usuario) Registrar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		usuario, err := u.service.Registrar(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, usuario, ""))
	}
}

func (u *Usuario) Modificar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			c.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
			return
		}
		if req.Apellido == "" {
			c.JSON(400, web.NewResponse(400, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Email == "" {
			c.JSON(400, web.NewResponse(400, nil, "El email del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			c.JSON(400, web.NewResponse(400, nil, "La edad es requerida"))
			return
		}
		if req.Altura == 0 {
			c.JSON(400, web.NewResponse(400, nil, "La altura es requerida"))
			return
		}
		if req.FechaCreacion == "" {
			c.JSON(400, web.NewResponse(400, nil, "La fecha de creacion del usuario es requerido"))
			return
		}

		u, err := u.service.Modificar(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, u, ""))
	}
}

func (c *Usuario) Eliminar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		err = c.service.Eliminar(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, id, ""))
	}
}

func (c *Usuario) ModificarAE() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad es requerida"))
			return
		}

		u, err := c.service.ModificarAE(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}
