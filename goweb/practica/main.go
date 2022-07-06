package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id            string  `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fechaCreacion"`
}

/* var users = []user{
	{
		Id:            "62c5b68a4f7187e949c458bb",
		Nombre:        "Tammi",
		Apellido:      "Prince",
		Email:         "tammiprince@pyramia.com",
		Edad:          36,
		Altura:        1.63,
		Activo:        false,
		FechaCreacion: "2016-07-12T09:58:39 +05:00",
	},
}  */

func getAll(c *gin.Context) {
	users := make([]user,0)
	raw, err := ioutil.ReadFile("./users.json")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error al cargar todos los usuarios",
		})
	}
	json.Unmarshal(raw, &users)
	c.JSON(200, users)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola" + " " + c.Query("nombre"),
		})
	})

	router.GET("/users", getAll)

	//puerto 8080
	router.Run()
}
