package handler

import (
	"fmt"
	"goweb/Clase3-2-WebServers/internal/usuarios"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Edad            int    `json:"edad"`
	Altura          int    `json:"altura"`
	Activo          bool   `json:"activo"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUser(user usuarios.Service) *Usuario {
	return &Usuario{
		service: user,
	}
}
func (user *Usuario) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		stringErrores := ""
		if req.Nombre == "" {
			stringErrores = stringErrores + "El nombre del usuario es requerido\n"
		}

		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Email == "" {
			stringErrores = stringErrores + "El email del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad es requerida\n"
		}
		if req.Altura == 0 {
			stringErrores = stringErrores + "La altura es requerida\n"
		}

		if len(stringErrores) != 0 {
			c.JSON(400, gin.H{"error": stringErrores})
			return
		}
		u, err := user.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, u)
	}
}

func (user *Usuario) UpdateSurnameAndAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		stringErrores := ""
		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad del usuario es requerida\n"
		}
		if len(stringErrores) != 0 {
			c.JSON(400, gin.H{"error": stringErrores})
			return
		}

		u, err := user.service.UpdateSurnameAndAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, u)
	}
}

func (u *Usuario) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		user, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, user)
	}
}

func (u *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		stringErrores := ""
		if req.Nombre == "" {
			stringErrores = stringErrores + "El nombre del usuario es requerido\n"
		}

		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Email == "" {
			stringErrores = stringErrores + "El email del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad es requerida\n"
		}
		if req.Altura == 0 {
			stringErrores = stringErrores + "La altura es requerida\n"
		}

		if len(stringErrores) != 0 {
			ctx.JSON(400, gin.H{"error": stringErrores})
			return
		}

		user, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, user)
	}
}
