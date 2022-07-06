package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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
		fmt.Println("error en lectura", err)
		ctx.JSON(404, err)
	}

	var users []user
	json.Unmarshal([]byte(data), &users)
	ctx.JSON(200, users)
}

func GetOne(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println("error en lectura", err)
		ctx.JSON(404, err)
	}

	ctx.JSON(200, data)

}

//siempre los filtrados por query

func main() {
	server := gin.Default()

	server.GET("/users", GetAll)
	server.GET("/one", GetOne)
	// c.JSON(200, gin.H{"msg": "Hola, Eimita"})

	server.Run()

}
