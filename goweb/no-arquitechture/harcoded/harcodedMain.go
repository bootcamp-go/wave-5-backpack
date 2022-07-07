package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id int
	Name string
	LastName string
	Email string
	Age int
	Height float32
	Active bool
	CreatedAt string
}

var (
	u1 = Usuario{
		Id: 1,
		Name: "Jorge",
		LastName: "Gonzalez",
		Email: "jorge@mail.com",
		Age: 30,
		Height: 2.02,
		Active: true,
		CreatedAt: "21/06/2",
	}

	u2 = Usuario{
		Id: 2,
        Name: "Romina",
        LastName: "Gutierrez",
        Email: "romi@mail.com",
        Age: 29,
        Height: 1.64,
        Active: false,
        CreatedAt: "03/10/20",
	}

)

var miSlice = []Usuario{u1,u2}



// trayendo el array harcodeado
func GetAll1(c *gin.Context){
	c.JSON(200, miSlice)
}

// Buscar por Id
func GetUserById1(c *gin.Context){
	
	id,_ := strconv.Atoi(c.Param("id"))
	
	for _,usuario := range miSlice{
		if usuario.Id == id {
			c.JSON(200,usuario)
			return
		}
	}
	c.JSON(404, "No se encontró ningún usuario con ese Id")
}


func main(){
	
	router := gin.Default()

	router.GET("/inicio", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hola Franco!",
		})
	})

	router.GET("/users1", GetAll1)
	router.GET("/users1/:id", GetUserById1) 

	router.Run(":8080")
}

