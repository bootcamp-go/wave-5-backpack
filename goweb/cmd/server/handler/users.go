package handler

import (
	"goweb/internal/users"
	"github.com/gin-gonic/gin"
	"strconv"
)

type request struct {
	Id int 				`json:"id"`
	Name string			`json:"name" binding:"required"`
	LastName string		`json:"lastname" binding:"required"`			
	Email string		`json:"email" binding:"required"`
	Age int				`json:"age"`
	Height float32		`json:"height"`
	Active bool			`json:"active"`
	CreatedAt string	`json:"createdat"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User{
	return &User{
		service: u,
	}
}

func (c *User) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		u, err := c.service.GetAllUsers()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id,_ := strconv.Atoi(ctx.Param("id"))
		u, err := c.service.GetUserById(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context){

		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		// traigo los datos del post y los guardo en una variable del tipo struct request que generé arriba
		var req request
		if err := ctx.Bind(&req); err !=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		newUser, err := c.service.StoreUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreatedAt)
		if err != nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, newUser)
	}
}