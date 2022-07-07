package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

func validateToken(c *gin.Context) bool {
	if token := c.GetHeader("token"); token != "123" {
		c.JSON(401, gin.H{
			"error": "Token Invalido",
		})
		return false
	}
	return true
}

func GetUsers(c *gin.Context) {
	var data, _ = os.ReadFile("./users.JSON")
	var u []User

	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"Saludo": "hola estas pasando por el userHandler",
		"data":   u,
	})
}

func GetUsersById(c *gin.Context) {
	var data, _ = os.ReadFile("./users.JSON")
	var u []User

	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	for key, _ := range u {
		if strconv.Itoa(u[key].Id) == c.Param("id") {
			c.JSON(200, gin.H{
				"saludo": "hola estas pasando por el GetuserByIdHandler",
				"data":   u[key],
			})
		}
	}
}

func CreateUser(c *gin.Context) {
	var userReq CreateUserRequest
	var user = User{}
	if validateToken(c) {
		if err := c.ShouldBindJSON(&userReq); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		userReq.Id = 3
		user.createUser(userReq)
		c.JSON(200, userReq)
	}
}
