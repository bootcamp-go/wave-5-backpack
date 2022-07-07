package main

import (
	"encoding/json"
	"fmt"

	"goweb/functions"
	"goweb/models"
	"goweb/services"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Saludar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola" + " " + c.Query("nombre"),
	})
}

func GetAll(c *gin.Context) {
	users, err := services.Read()
	if err != nil {
		c.JSON(500, gin.H{"Error:": err.Error()})
		return
	}
	c.JSON(200, users)
}

func GetUserById(ctx *gin.Context) {
	users, err := services.Read()
	id := ctx.Param("id")

	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}

	for _, user := range users {
		if user.Id == id {
			ctx.JSON(200, user)
			return
		}
	}
	ctx.JSON(404, gin.H{"Error:": "User not found"})
}

var users, _ = services.Read()

func CreateUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			ctx.JSON(400, functions.ValidateErrors(fieldError.Field(), fieldError))
		}
		return
	}
	newUser.Id = functions.IdRandom(10)

	users = append(users, newUser)

	file, _ := json.MarshalIndent(users, "", "\t")
	fmt.Println(string(file))

	_ = ioutil.WriteFile("./users.json", file, 0644)

	ctx.JSON(200, newUser)
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/saludar", Saludar)
	user := router.Group("/users")
	{
		user.GET("/", GetAll)
		user.GET("/:id", GetUserById)
		user.POST("/", CreateUser)
	}

	//puerto 8080
	router.Run()
}
