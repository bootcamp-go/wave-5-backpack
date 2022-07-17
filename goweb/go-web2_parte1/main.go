package main

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	Id int `json:id`
	Nombre string `json:nombre`
	Color string	`json:color`
	Precio int	`json:precio`
	Stock int	`json:stock`
	Codigo string	`json:codigo`
	Publicado bool	`json:publicado`
	Fecha string	`json:fecha`
}

var products []Request
var lastId int

func main()  {

	router := gin.Default()

	gr := router.Group("/productos")
	gr.POST("/save", Save())

	router.Run()
}

func Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123"{
			c.JSON(401, gin.H{"error": "token invalido"})
			return
		}
		var productos Request
		if err := c.ShouldBindJSON(&productos); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if  productos.Nombre == ""{
			c.JSON(400, gin.H{"Error": "el campo Nombre es requerido"})
			return
		}
		if  productos.Color == ""{
			c.JSON(400, gin.H{"Error": "el campo Color es requerido"})
			return
		}
		if  productos.Precio == 0{
			c.JSON(400, gin.H{"Error": "el campo Precio es requerido"})
			return
		}
		if  productos.Stock == 0{
			c.JSON(400, gin.H{"Error": "el campo stock es requerido"})
			return
		}
		if  productos.Codigo == ""{
			c.JSON(400, gin.H{"Error": "el campo codigo es requerido"})
			return
		}
		if  productos.Publicado != true && productos.Publicado != false{
			c.JSON(400, gin.H{"Error": "el campo publicado es requerido"})
			return
		}
		if  productos.Fecha == ""{
			c.JSON(400, gin.H{"Error": "el campo fecha es requerido"})
			return
		}

		lastId++
		productos.Id = lastId
		products = append(products ,productos)
		c.JSON(200, productos)
	}
}
