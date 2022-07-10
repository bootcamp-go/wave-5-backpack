package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type users struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Edad     int     `json:"edad" binding:"required"`
	Altura   float64 `json:"altura" binding:"required"`
	Activo   bool    `json:"activo" binding:"required"`
}

var usuarios []users
var U = users{}
var lastID int

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)

	pr := router.Group("/usuarios")
	{
		pr.POST("/", GuardarWithAuthorize())
	}
	router.Run()

}

func GetAll(c *gin.Context) {

	data, err := os.ReadFile("./usuarios.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se abrio el archivo.")
	if err = json.Unmarshal([]byte(data), &U); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, U)
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req users
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": fmt.Errorf("Error"),
			})
			return
		}

		lastID++

		req.Id = lastID
		usuarios = append(usuarios, req)

		c.JSON(200, req)
	}
}

func GuardarWithAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req users
		// Validar token
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		// Si el token fue valido, avanzo
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.Id = lastID

		usuarios = append(usuarios, req)

		c.JSON(200, usuarios)
	}

}
