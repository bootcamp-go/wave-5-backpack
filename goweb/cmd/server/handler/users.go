package handler

import (
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name       string  `json: "name"`
	Lastname   string  `json: "lastname"`
	Email      string  `json: "email"`
	Age        int     `json: "age"`
	Height     float32 `json: "height"`
	Active     bool    `json: "active"`
	DoCreation string  `json: "doCreation"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "lalala" {
			ctx.JSON(401, gin.H{
				"ERROR": "Invalid token",
			})
			return
		}
		allUsers, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		ctx.JSON(200, allUsers)
	}
}

func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "lalala" {
			ctx.JSON(401, gin.H{
				"ERROR": "Invalid token",
			})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		newUser, err := c.service.StoreUser(req.Name, req.Lastname, req.Email, req.Age, req.Height, req.Active, req.DoCreation)
		if err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		ctx.JSON(200, newUser)
	}
}
