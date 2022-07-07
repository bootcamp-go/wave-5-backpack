package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id        int    `form:"id" json:"id"`
	Nombre    string `form:"nombre" json:"nombre"`
	Color     string `form:"color" json:"color"`
	Precio    int    `form:"precio" json:"precio"`
	Stock     int    `form:"stock" json:"stock"`
	Codigo    string `form:"codigo" json:"codigo"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha"`
}

func getJson() []producto {
	jsonFile, err := os.Open("./products.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var productos []producto
	json.Unmarshal([]byte(byteValue), &productos)
	// ctx.JSON(200, productos)
	defer jsonFile.Close()
	return productos
}

func getAll(ctx *gin.Context) {

	// Filtrando con todos los campos
	var struct_producto producto

	if err := ctx.ShouldBindQuery(&struct_producto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapKeys := make(map[string]interface{})
	mapKeys["string"] = ""
	mapKeys["int"] = 0
	mapKeys["bool"] = false

	mapProduct := jsonToMap(struct_producto)[0]
	var keysList = []string{}
	for key, value := range mapProduct {
		if value != mapKeys[reflect.TypeOf(value).String()] {
			keysList = append(keysList, key)
		}
	}

	productos := jsonToMap(getJson()...)
	var filtered_products []map[string]interface{}
	var filtered_products_empty []map[string]interface{}

	for _, key := range keysList {
		for _, p := range productos {
			if p[key] == mapProduct[key] {
				filtered_products = append(filtered_products, p)
			}
		}
		fmt.Println(filtered_products)
		productos = filtered_products
		filtered_products = filtered_products_empty
	}

	ctx.JSON(200, productos)

	// Filtrando solo por el ID
	/*
		productos := getJson()
		var list []producto
		for _, p := range productos {
			if p.Id == struct_producto.Id {
				list = append(list, p)
			}
		}
		ctx.JSON(200, list) */

}

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
		productoMap["fecha"] = p.Fecha

		productoMaps = append(productoMaps, productoMap)
		productoMap = make(map[string]interface{})
	}

	return productoMaps
}

func getById(ctx *gin.Context) {
	productos := getJson()
	id, _ := strconv.Atoi(ctx.Param("id"))
	var producto_id producto
	for _, p := range productos {
		if id == p.Id {
			producto_id = p
			break
		}
	}
	ctx.JSON(200, producto_id)
}

func main() {
	router := gin.Default()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	msg := "Hola " + name
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": msg,
		})
	})

	router.GET("/productos", getAll)
	router.GET("/productos/:id", getById)

	router.Run()
}
