package handler

import (
	"time"

	"ejercicioTM/internal/users"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		usuario, err := c.service.GetOne(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		err := c.service.Delete(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Status(204)
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Update(ctx.Param("id"), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
