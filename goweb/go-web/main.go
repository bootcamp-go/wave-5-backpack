package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required`
	Surname string `json:"surname" binding:"required`
	Email string `json:"email" binding:"required`
	Age int `json:"age" binding:"required`
	Active bool `json:"active" binding:"required`
	Created string `json:"created" binding:"required`
}

var Request []User

func main(){
	router := gin.Default()
	
	router.GET("/saludo", say)
	users := router.Group("/users")
	{
		users.GET("/", get)
		users.GET("/:id", getOne)
		users.GET("/", filter)
		users.POST("/user", CrearEntidad)
	}
	
	

	router.Run(":8080")
}


func say(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "hi nahu!",		})
}

func get(ctx *gin.Context){
	u, err := read()
	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
	}
	ctx.JSON(200, gin.H{
		"message": 	u,
	})
}

func getOne(ctx *gin.Context){
	u, err := read()
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}

	for _, user := range u {
		if user.Id == id {
			ctx.JSON(200, user)
			return
		}
	}
	ctx.JSON(404, gin.H{"Error:": "User not found"})
	
}

func read()([]User, error){
	jsonFile, _ := os.Open("usuarios.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var u []User
	err := json.Unmarshal(byteValue, &u)
	if err != nil {
		return []User{}, errors.New("tuvimos un problema en la serialización")
	}
	return u, nil
}

func filter(ctx *gin.Context){
	name := ctx.Query("name")
	surname := ctx.Query("surname")
	age, _ := strconv.Atoi(ctx.Query("age"))
	users, err := read()
	var res []User

	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
	}
	for _, value := range users {
		if value.Name == name{
			res = append(res, value)
		}
		if value.Age == age {
			res = append(res, value)
		}
		if value.Surname == surname {
			res = append(res, value)
		}
	}
	if len(res) > 0 {
		ctx.JSON(200, res)
	}
	ctx.JSON(404, gin.H{"Error": errors.New("no se encontraron resultados")})
}

func CrearEntidad(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(401, gin.H{
			"error": "token invalido",
		})
		return
	}

	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, valError := range errs {
			if valError.Tag() == "required" {
				c.JSON(400, gin.H{
					"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
				})
				return
			}
		}

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Id = Maximo(Request) + 1
	Request = append(Request, req)
	c.JSON(201, req)
}

func Maximo(req []User) int {
	maximo := 0
	for _, user := range req {
		if maximo < user.Id {
			maximo = user.Id
		}
	}

	return maximo
}


/*
PASARLE PARAMETROS A UN CONTROLADOR
func Guardar(parametro) gin.HandlerFunc {
	return func(c *gin.Context){
		var req request
		etc
		etc
		etc
	}
}
*/