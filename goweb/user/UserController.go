package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
func CreateUser(u User) gin.HandlerFunc {
	var user CreateUserRequest
	return func(c *gin.Context) {
		if !validateToken(c) {
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.ActualTag())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.ActualTag())
					}
				}
				c.JSON(404, result)
				return
			}
		}
		//TODO: Llamar al servicio interno createUser
		u.createUser(user)
		c.JSON(200, user)
	}
}
