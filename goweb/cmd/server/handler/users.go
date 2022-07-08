package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string  `json:"name" binding:"required"`
	LastName     string  `json:"last_name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creation_date" binding:"required"`
}

type patchRequest struct {
	LastName string `json:"last_name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(p users.Service) *User {
	return &User{
		service: p,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		u, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
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
		u, err := u.service.Store(req.Age, req.Name, req.LastName, req.Email, req.CreationDate, req.Height, req.Active)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
			return
		}

		if req.LastName == "" {
			ctx.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}

		// validar email con regex
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "El email del usuario es requerido"})
			return
		}

		if req.Age < 1 {
			ctx.JSON(400, gin.H{"error": "La edad del usuario debe ser mayor a 0."})
			return
		}

		if req.Height <= 0 {
			ctx.JSON(400, gin.H{"error": "La altura del usuario es requerida"})
			return
		}

		// preguntar, es raro
		if !req.Active {
			ctx.JSON(400, gin.H{"error": "Es requerido saber si es activo el usuario"})
			return
		}

		if req.CreationDate == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creacion del usuario es requerida"})
			return
		}

		p, err := u.service.Update(id, req.Age, req.Name, req.LastName, req.Email, req.CreationDate, req.Height, req.Active)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func (u *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req patchRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.LastName == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Age < 1 {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		u, err := u.service.UpdateLastNameAndAge(id, req.Age, req.LastName)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = u.service.Delete(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
