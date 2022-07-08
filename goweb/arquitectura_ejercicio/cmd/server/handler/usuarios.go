package handler

import (
	"net/http"
	"strconv"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/usuarios"
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
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "El token es inválido",
			})
			return
		}
		usuarios, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, usuarios)
	}
}

func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := c.service.Store(req.Age, req.Name, req.LastName, req.Email, req.Estatura)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req domain.Usuario
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := c.service.Update(id, req.Age, req.Names, req.LastName, req.Email, req.DateCreated, req.Estatura, req.IsActivo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *Usuario) PatchLastNameAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req LastNameAgePatchRequest
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := c.service.UpdateLastNameAndAge(id, req.Age, req.LastName)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
