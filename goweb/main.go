package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id         int
	Name       string
	Lastname   string
	Email      string
	Age        int
	Height     float64
	Active     bool
	DoCreation string
}

func GetAll(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	var users []user
	json.Unmarshal([]byte(data), &users)
	ctx.JSON(200, users)
}

func GetById(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")
	var users []user
	json.Unmarshal([]byte(data), &users)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	id := ctx.Param("id")

	for _, user := range users {
		if fmt.Sprint(user.Id) == id{
			ctx.JSON(200, user)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": err.Error()})

}

//siempre los filtrados por query

func main() {
	server := gin.Default()

	users := server.Group("/users")
	{
		users.GET("/", GetAll)
		users.GET("/:id", GetById)
	}

	// c.JSON(200, gin.H{"msg": "Hola, Eimita"})

	server.Run(":8080")

}
