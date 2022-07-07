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
	Id       int    `json:"id" binding:"-"`
	Codigo   string `json:"codigo" binding:"required"`
	Moneda   string `json:"moneda" binding:"required"`
	Monto    int    `json:"monto" binding:"required"`
	Emisor   string `json:"emisor" binding:"required"`
	Receptor string `json:"receptor" binding:"required"`
	Fecha    string `json:"fecha" binding:"required"`
}

var transactions []transacciones
var lastID int = 9

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

func postTrans() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "12345" {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		var req transacciones
		if err := c.ShouldBindJSON(&req); err != nil {
			if req.Codigo == "" {
				c.JSON(404, gin.H{"error": "El campo código es requerido"})
			}
			if req.Moneda == "" {
				c.JSON(404, gin.H{"error": "El campo moneda es requerido"})
			}
			if req.Monto == 0 {
				c.JSON(404, gin.H{"error": "El campo monto es requerido"})
			}
			if req.Emisor == "" {
				c.JSON(404, gin.H{"error": "El campo emisor es requerido"})
			}
			if req.Receptor == "" {
				c.JSON(404, gin.H{"error": "El campo receptor es requerido"})
			}
			if req.Fecha == "" {
				c.JSON(404, gin.H{"error": "El campo fecha es requerido"})
			}
			// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// fmt.Println(string(err.Error()))
			return
		}

		maxId := 0

		for i, value := range transactions {
			if value.Id > maxId {
				maxId = value.Id
				fmt.Printf("Entro %d \n", i)
			}
		}

		lastID = maxId + 1
		req.Id = lastID
		transactions = append(transactions, req)
		c.JSON(http.StatusAccepted, transactions)
	}
}

func main() {
	data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		fmt.Errorf("Se produjo un error al leer el archivo\n")
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		fmt.Errorf("Se produjo un error al traducir a Go\n")
	}

	router := gin.Default()

	router.GET("/nombre", func(c *gin.Context) {
		name := c.Query("name")
		c.String(200, "Mi nombre es "+name)
	})

	router.GET("/transacciones", getAll)
	router.GET("/transacciones/filtros", getByQuery)
	router.GET("/transacciones/:id", getById)

	router.POST("/transacciones", postTrans())

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
