package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id           int     `json:"-"`
	FirstName    string  `json:"firstName" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

var users []User
var lastId int

func main() {
	router := gin.Default()
	router.POST("/users", CreateUser())
	router.Run()
}
func validateError(er error) {
	log.Println(er)
}
func validateToken(c *gin.Context) bool {
	if token := c.GetHeader("token"); token != "123" {
		c.JSON(401, gin.H{
			"error": "No tiene permisos para realizar la peticion solicitada",
		})
		return false
	}
	return true
}

func CreateUser() gin.HandlerFunc {
	var user User
	return func(c *gin.Context) {
		if !validateToken(c) {
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			validateError(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastId++
		user.Id = lastId
		users = append(users, user)
		c.JSON(200, users)
	}
}
