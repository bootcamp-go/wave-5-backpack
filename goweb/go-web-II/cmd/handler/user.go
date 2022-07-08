package handler

import (
	"fmt"
	"goweb/go-web-II/internal/products"
	"net/http"
	"strconv"

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
ac
á va la data que recibimos del body
*/

type User struct {
	service products.Service
}


func NewUser(s products.Service)*User{
	return &User{service: s}
}

func tokenValidator(c *gin.Context){
	token := c.GetHeader("token")
	if token != "123"{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
		return
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func (c *gin.Context){
		tokenValidator(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = u.service.Delete(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("el producto %d ha sido eliminado", id)})
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("token")
		if token != "123"{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Age == 0 {
			c.JSON(400, gin.H{"error": "la edad es necesaria"})
		}
		if req.Created == "" {
			c.JSON(400, gin.H{"error": "la fecha de creacion es necesaria"})
		}
		if req.Email == "" {
			c.JSON(400, gin.H{"error": "el email es necesario"})
		}
		if req.Name == "" {
			c.JSON(400, gin.H{"error": "el nombre es necesario"})
		}
		if req.Surname == "" {
			c.JSON(400, gin.H{"error": "el apellido es necesario"})
		}
		u, err := u.service.Update(id, req.Age, req.Name, req.Surname, req.Email, req.Created, req.Active)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, u)
	}
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