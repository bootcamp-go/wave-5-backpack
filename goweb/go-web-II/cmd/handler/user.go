package handler

import (
	"goweb/go-web-II/internal/products"
	"net/http"

	"github.com/gin-gonic/gin"
)


type request struct {
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age int `json:"age" binding:"required"`
	Active bool `json:"active"`
	Created string `json:"created" binding:"required"`
}
/*
acá va la data que recibimos del body
*/

type User struct {
	service products.Service
}


func NewUser(s products.Service)*User{
	return &User{service: s}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func (c *gin.Context){
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}
		users, err := u.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token inválido"})
		}
		if len(users) <= 0{
			c.JSON(http.StatusOK, gin.H{
				"message": "no se encontraron usuarios"})
		}
		c.JSON(http.StatusOK, gin.H{"usuarios":users})
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func (c *gin.Context){
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req)
		err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := u.service.Store(req.Age, req.Name, req.Surname, req.Email, req.Created, req.Active)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}