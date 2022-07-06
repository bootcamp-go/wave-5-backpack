package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

// estructura a recibir en la URL
type user struct {
	Name      string `form:"name"`
	Last_name string `form:"last_name"`
}

// estructura escogida para el ejercicio
type producto struct {
	Id            int
	Nombre, Color string
	Precio, Stock int
	Codigo        string
	Publicado     bool
	FechaCreacion string
}

func main() {
	router := gin.Default()

	saludoGet(router)
	productosGet(router)

	router.Run()
}

// funcion que recibe los datos ingresados por url
// y los devuelve en un mensaje de saludo. Si no
// se reciben datos, saluda a un usuario visitante
func saludoGet(router *gin.Engine) {

	router.GET("/saludo", func(c *gin.Context) {
		var userObj user
		var emptyUser user

		if err := c.ShouldBindQuery(&userObj); err == nil {
			fmt.Println("User obj: ", userObj)
			if userObj == emptyUser {
				userObj.Name = "usuario"
				userObj.Last_name = "visitante"
			}
		} else {
			fmt.Println("Error: ", err)
		}

		c.JSON(200, gin.H{
			"message": "Hola " + userObj.Name + " " + userObj.Last_name,
		})
	})
}

// Recibir productos de una lista creada manualmente y de un archivo JSON
func productosGet(router *gin.Engine) {

	router.GET("/productos", getAll)

	router.GET("/productos_archivo", getAllFile)
}

// Crea una lista de productos por defecto y la envía como respuesta
func getAll(c *gin.Context) {
	var productos = []producto{
		{Id: 4, Nombre: "Tenedor", Color: "Rojo", Precio: 10, Stock: 50, Codigo: "C000", Publicado: true, FechaCreacion: "10-08-2005"},
		{Id: 5, Nombre: "Cuchillo", Color: "Verde", Precio: 8, Codigo: "C001", Publicado: false, FechaCreacion: "12-03-2011"},
		{Id: 6, Nombre: "Cuchara", Color: "Morado", Precio: 5, Codigo: "C002", Publicado: true, FechaCreacion: "14-1-2016"},
	}

	c.JSON(200, productos)
}

// Abre el archivo products.json, lo lee en datos de bytes y usa
// Unmarshal para convertirlos en una lista de productos que
// posteriormente envia como respuesta
func getAllFile(c *gin.Context) {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}

	var productos []producto
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &productos)
	c.JSON(200, productos)
	jsonFile.Close()
}
