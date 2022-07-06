package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id, Age                               int
	FirstName, LastName, Email, CreatedAt string
	Height                                float64
	Active                                bool
}

func ReadJSONFile(fileName string) ([]User, error) {
	data, err := os.ReadFile(fileName)
	users := []User{}
	if err == nil {
		err = json.Unmarshal(data, &users)
	}
	return users, err
}

func GetAll(c *gin.Context) {
	users, err := ReadJSONFile("users.json")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		fmt.Printf("Error: %v\n", err)
		return
	}
	c.JSON(200, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", GetAll)
	router.Run()
}
