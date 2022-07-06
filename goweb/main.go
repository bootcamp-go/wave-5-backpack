package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transacciones struct {
	Id       int    `json: "id"`
	Codigo   string `json: "codigo"`
	Moneda   string `json: "moneda"`
	Monto    int    `json: "monto`
	Emisor   string `json: "emisor"`
	Receptor string `json: "receptor"`
	Fecha    string `json: "fecha"`
}

var transactions []transacciones

func getAll(c *gin.Context) {
	c.JSON(200, gin.H{"data": transactions})
}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	for _, trans := range transactions {
		if trans.Id == id {
			c.JSON(http.StatusAccepted, trans)
		}
	}
	c.JSON(404, gin.H{"error": "ID no existente"})
}

func getByQuery(c *gin.Context) {
	monto, _ := strconv.Atoi(c.Query("monto"))
	for _, t := range transactions {
		if t.Codigo == c.Query("codigo") {
			c.JSON(http.StatusAccepted, t)
		}
		if t.Moneda == c.Query("moneda") {
			c.JSON(http.StatusAccepted, t)
		}
		if t.Monto == monto {
			c.JSON(http.StatusAccepted, t)
		}
		if t.Emisor == c.Query("emisor") {
			c.JSON(http.StatusAccepted, t)
		}
		if t.Receptor == c.Query("receptor") {
			c.JSON(http.StatusAccepted, t)
		}
		if t.Fecha == c.Query("fecha") {
			c.JSON(http.StatusAccepted, t)
		}
	}
}

func main() {
	data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		fmt.Errorf("Se produjo un error al leer el archivo")
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		fmt.Errorf("Se produjo un error al traducir a Go")
	}

	router := gin.Default()

	router.GET("/nombre", func(c *gin.Context) {
		name := c.Query("name")
		c.String(200, "Mi nombre es "+name)
	})

	router.GET("/transacciones", getAll)
	router.GET("/transacciones/filtros", getByQuery)
	router.GET("/transacciones/:id", getById)

	router.Run(":8080")
}

//Ejercicio 2
//Crea dentro de la carpeta go-web un archivo llamado main.go
//Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
//Pegale al endpoint para corroborar que la respuesta sea la correcta.

//Ejercicio 3
//Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
//Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
//Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
//Genera un handler para el endpoint llamado “GetAll”.
//Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
