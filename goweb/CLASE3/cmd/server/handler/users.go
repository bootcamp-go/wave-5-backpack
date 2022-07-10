package handler

import (
	"clase2_parte2/internal/users"
	"net/http"
	"os"

	"fmt"
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		usuario, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (c *Users) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := c.service.Store(req.Nombre, req.Apellido, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (u *Users) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El id no es valido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"Erro": "Nombre requerido"})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"Erro": "Apellido requerido"})
			return
		}
		if req.Edad <= 0 {
			ctx.JSON(400, gin.H{"Erro": "Edad requerida"})
			return
		}
		if req.Altura <= 0 {
			ctx.JSON(400, gin.H{"Erro": "Altura requerida"})
			return
		}

		u, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(400, gin.H{"erro": err.Error()})
			return
		}
		ctx.JSON(200, u)

	}
}

func (u *Users) UpdateApellidoAndEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El id no es valido"})
			return
		}
		type request struct {
			Apellido string `json:"apellido" binding:"required"`
			Edad     int    `json:"edad" binding:"required"`
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"Erro": "Apellido requerida"})
			return
		}
		if req.Edad <= 0 {
			ctx.JSON(400, gin.H{"Erro": "Edad erronea"})
			return
		}

		u, err := u.service.UpdateApellidoAndEdad(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)

	}
}

func (u *Users) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(400, gin.H{"error": "Token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
