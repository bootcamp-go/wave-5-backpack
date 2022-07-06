package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id, Edad                                int
	Nombre, Apellido, Email, Fecha_creacion string
	Altura                                  float64
	Activo                                  bool
}

func main() {
	router := gin.Default()

	// router.GET("/HolaNombre", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"mesage": "Hola Jose",
	// 	})
	// })
	router.GET("/GetAll", func(ctx *gin.Context) {
		var u []Usuarios
		j, _ := os.ReadFile("./usuarios.json")
		if err := json.Unmarshal(j, &u); err != nil {
			log.Println(string(j))
			log.Fatal(err)
		}
		ctx.JSON(200, gin.H{
			"usuario": u,
		})
	})
	router.Run()
}
