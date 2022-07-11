package handler

import (
	"ejercicio-2-3/internal/usuarios"
	"fmt"
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		usuarios, err := u.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token inválido"}) // 500
			return
		}

		if len(usuarios) <= 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No se encontraron usuarios."}) // 500
			return
		}

		c.JSON(http.StatusOK, gin.H{"usuarios": usuarios})
	}
}

func (u *Usuario) Registrar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		usuario, err := u.service.Registrar(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //400
			return
		}

		c.JSON(http.StatusOK, gin.H{"usuario": usuario})
	}
}

func (c *Usuario) Modificar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "El email del usuario es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "La edad es requerida"})
			return
		}
		if req.Altura == 0 {
			ctx.JSON(400, gin.H{"error": "La altura es requerida"})
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creacion del usuario es requerido"})
			return
		}

		u, err := c.service.Modificar(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *Usuario) Eliminar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Eliminar(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}

func (c *Usuario) ModificarAE() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "La edad es requerida"})
			return
		}

		u, err := c.service.ModificarAE(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}
