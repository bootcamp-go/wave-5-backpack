package main

import "github.com/gin-gonic/gin"

type usuarios struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_de_creacion"`
}

func GetAll(ctx *gin.Context) {
	users := []usuarios{{}, {}}
	ctx.JSON(200, gin.H{
		"usuarios": users,
		"abc":      "prueba",
	})
}

func main() {
	router := gin.Default()
	router.GET("hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Andres!",
		})
	})
	router.GET("usuarios", GetAll)
	router.Run()
}
