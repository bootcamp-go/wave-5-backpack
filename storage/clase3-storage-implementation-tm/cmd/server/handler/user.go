package handler

import (
	"clase3-storage-implementation-tm/internal/users"

	"github.com/gin-gonic/gin"
)

type request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

// User ...
type User struct {
	service users.Service
}

// NewUser ...
func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

// GetOne ...
func (c *User) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		p, err := c.service.GetOne(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

// Store ...
func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Firstname, req.Lastname, req.Username, req.Email)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
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

		p, err := c.service.Update(ctx.Param("id"), req.Firstname,
			req.Lastname, req.Username, req.Email)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

// Delete ...
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
