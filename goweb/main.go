package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

func main() {
	router := gin.Default()

	router.GET("/usuarios", GetAll)

	router.Run()
}

func GetAll(c *gin.Context) {
	var u Users
	jsonData, _ := os.ReadFile("usuarios.json")

	// fmt.Println(jsonData)

	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(u)

	c.JSON(200, gin.H{
		"message": u.Usuarios,
	})
}
