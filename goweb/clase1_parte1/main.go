package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"log"
	"os"
)

type usuario struct {
	Id 				int		`json:"id"`
	Nombre 			string	`json:"nombre"`
	Apellido 		string	`json:"apellido"`
	Email 			string	`json:"email"`
	Edad 			int		`json:"edad"`
	Altura 			float64	`json:"altura"`
	Activo			bool	`json:"activo"`
	FechaCreacion 	string	`json:"fecha_creacion"`
}

func GetAll() gin.HandlerFunc {

	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []usuario
	if err := json.Unmarshal([]byte(jsonUsers), &users); err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		c.JSON(200, users)
		// c.JSON(200, gin.H {
		// 	"usuarios": users,
		// })
	}
}

func main() {
	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"saludo": "Hola Diana",
		})
	})

	router.GET("/usuarios", GetAll())
	router.Run()
}


