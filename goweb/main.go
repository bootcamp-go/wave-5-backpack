package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type usuarios struct {
	Id                      int `json:"id" binding:"required"`
	Nombre, Apellido, Email string
	Edad                    int
	Altura                  float64
	Activo                  bool
	Fecha                   time.Time
}

func main() {
	//router con logger y recover
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		//respuesta tipo JSON
		c.JSON(200, gin.H{
			"message": "Hola Luz! :D",
		})
	})

	routerDos := gin.Default()

	routerDos.GET("/usuarios", getAll)
	// Corriendo servidor sobre el puerto 8080
	router.Run()
	routerDos.Run(":8081")
}

func getAll(c *gin.Context) {
	users := []usuarios{
		{
			Id:       1,
			Nombre:   "Luz",
			Apellido: "Lucumí",
			Email:    "luz.lucumi@hotmail.com",
			Edad:     26,
			Altura:   1.65,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       2,
			Nombre:   "Luber",
			Apellido: "Lucumí",
			Email:    "luber.lucumi@hotmail.com",
			Edad:     61,
			Altura:   1.82,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       3,
			Nombre:   "Martha",
			Apellido: "Hernández",
			Email:    "martha@hotmail.com",
			Edad:     60,
			Altura:   1.60,
			Activo:   true,
			Fecha:    time.Now(),
		},
	}
	c.JSON(200, gin.H{"usuarios": users})
}
