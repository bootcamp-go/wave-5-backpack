package handler

import (
	"goweb/internal/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SECRET_TOKEN = "123456"

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": users,
		})
	}
}
