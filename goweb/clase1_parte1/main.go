package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// estructura a recibir en la URL
type user struct {
	Name      string `form:"name"`
	Last_name string `form:"last_name"`
}

// estructura escogida para el ejercicio
type producto struct {
	Id            int    `form:"id" json:"id"`
	Nombre        string `form:"nombre" json:"nombre"`
	Color         string `form:"color" json:"color"`
	Precio        int    `form:"precio" json:"precio"`
	Stock         int    `form:"stock" json:"stock"`
	Codigo        string `form:"codigo" json:"codigo"`
	Publicado     bool   `form:"publicado" json:"publicado"`
	FechaCreacion string `form:"fecha_creacion" json:"fecha_creacion"`
}

func main() {
	router := gin.Default()

	saludoGet(router)
	productosGet(router)

	router.Run()
}

// Recibe los datos ingresados por url y los devuelve en un mensaje
// de saludo. Si no se reciben datos, saluda a un usuario visitante
func saludoGet(router *gin.Engine) {

	router.GET("/saludo", func(c *gin.Context) {
		var userObj user
		var emptyUser user

		if err := c.ShouldBindQuery(&userObj); err == nil {
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

	productos_archivo := router.Group("/productos_archivo")
	{
		productos_archivo.GET("/", getAllFile)
		productos_archivo.GET("/:id", getAllFileID)
	}
	// router.GET("/productos_archivo", getAllFile)
	// router.GET("/productos_archivo/:id", getAllFileID)
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

// Filtra los productos por los parametros solicitados
func getAllFile(c *gin.Context) {

	var product producto
	if err := c.ShouldBindQuery(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapKeys := make(map[string]interface{})
	mapKeys["string"] = ""
	mapKeys["int"] = 0
	mapKeys["bool"] = false

	mapProduct := jsonToMap(product)[0]
	var keysList = []string{}
	for key, value := range mapProduct {
		if value != mapKeys[reflect.TypeOf(value).String()] {
			keysList = append(keysList, key)
		}
	}

	productos := jsonToMap(readJSON()...)
	var filtered_products []map[string]interface{}
	var filtered_products_empty []map[string]interface{}

	for _, key := range keysList {
		for _, p := range productos {
			if p[key] == mapProduct[key] {
				filtered_products = append(filtered_products, p)
			}
		}
		productos = filtered_products
		filtered_products = filtered_products_empty
	}

	c.JSON(200, productos)

}

// Obtiene la lista de productos del archivo products.json
func readJSON() []producto {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}

	var productos []producto
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &productos)
	jsonFile.Close()
	return productos
}

// Filtra los productos por el id solicitado
func getAllFileID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productos := readJSON()
	var filtered_product producto
	for _, p := range productos {
		if p.Id == id && id > 0 {
			filtered_product = p
			break
		}
	}
	c.JSON(200, filtered_product)
}

// Convierte una lista de struct productos a map[string]interface
func jsonToMap(productos ...producto) []map[string]interface{} {
	productoMap := make(map[string]interface{})
	var productoMaps []map[string]interface{}

	for _, p := range productos {
		productoMap["id"] = p.Id
		productoMap["nombre"] = p.Nombre
		productoMap["color"] = p.Color
		productoMap["precio"] = p.Precio
		productoMap["stock"] = p.Stock
		productoMap["codigo"] = p.Codigo
		productoMap["publicado"] = p.Publicado
		productoMap["fecha_creacion"] = p.FechaCreacion

		productoMaps = append(productoMaps, productoMap)
		productoMap = make(map[string]interface{})
	}

	return productoMaps
}
