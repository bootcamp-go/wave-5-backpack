package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Edad            int    `json:"edad"`
	Altura          int    `json:"altura"`
	Activo          bool   `json:"activo"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}

func HandlerGetAll(c *gin.Context) {
	/*var users = []usuario{
		{id: 1, nombre: "Leonardo", apellido: "Da Vinci", email: "leodavinci@gmail.com", edad: 504, altura: 163, activo: false, fechaDeCreacion: "03/09/1540"},
		{id: 2, nombre: "Salvador", apellido: "Dali", email: "sdali@gmail.com", edad: 120, altura: 172, activo: false, fechaDeCreacion: "03/09/1940"},
		{id: 3, nombre: "Pablo", apellido: "Picasso", email: "ppicasso@gmail.com", edad: 115, altura: 160, activo: true, fechaDeCreacion: "03/09/1980"},
	}
	fmt.Println(users)
	*/
	var listaUsuarios []usuario
	jsonData, _ := os.ReadFile("./usuarios.json")
	fmt.Println(string(jsonData))
	if err := json.Unmarshal((jsonData), &listaUsuarios); err != nil {
		log.Fatal(err)
	}
	datos, err := json.Marshal(listaUsuarios)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, string(datos))
}

func main() {

	//fmt.Println(listaUsuarios)
	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "German Rodriguez!!!"})
	})
	/*
	 */
	router.GET("/usuarios", HandlerGetAll)

	router.Run(":8080")
}
