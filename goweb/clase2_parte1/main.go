package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creacion"`
}

var products []Product
var lastId int

func main() {
	router := gin.Default()
	router.POST("/products", CreateProduct())
	router.Run()
}

func validateToken(c *gin.Context) bool {
	if token := c.GetHeader("token"); token != "123" {
		c.JSON(401, gin.H{
			"error": "No tiene permisos para realizar la peticion solicitada",
		})
		return false
	}
	return true
}

func CreateProduct() gin.HandlerFunc {
	var product Product
	return func(c *gin.Context) {
		if !validateToken(c) {
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var errStruct []string
		if product.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if product.Color == "" {
			errStruct = append(errStruct, "color")
		}
		if product.Precio <= 0 {
			errStruct = append(errStruct, "precio")
		}
		if product.Stock < 0 {
			errStruct = append(errStruct, "stock")
		}
		if product.Codigo == "" {
			errStruct = append(errStruct, "codigo")
		}
		if product.FechaCreacion == "" {
			errStruct = append(errStruct, "fecha_de_creacion")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Los campos %s son requeridos.", strings.Join(errStruct, ", "))})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("El campo %s es requerido.", errStruct[0])})
				return
			}
		}
		// {
		// 	"id": 1,
		// 	"nombre": "",
		// 	"color": "",
		// 	"precio": 0,
		// 	"stock": 0,
		// 	"codigo": "",
		// 	"publicado": false,
		// 	"fecha_de_creacion": ""
		// }
		lastId++
		product.Id = lastId
		products = append(products, product)
		c.JSON(200, product)
	}
}
