package handler

import (
	"net/http"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

// Estructura de rerquest
type request struct {
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
}

type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{
		service: s,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "449d451b-f411-4dc8-aefb-d8a33c723ffa" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "449d451b-f411-4dc8-aefb-d8a33c723ffa" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}
