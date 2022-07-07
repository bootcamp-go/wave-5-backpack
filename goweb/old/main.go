package main

import (
	"errors"
	"fmt"
	"goweb/models"
	"goweb/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Clase 1 Ejercicio 1 Parte 1
func GetAll(ctx *gin.Context) {
	products, err := services.ReadAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Ejercicio 1 parte 2
	var filtrados []models.Product
	for _, product := range products {

		var idVerif, nombreVerif, colorVerif, precioVerif, stockVerif, codigoVerif, publicadoVerif, fechaCreacionVerif bool

		idVerif = ctx.Query("id") != "" && ctx.Query("id") != fmt.Sprintf("%d", product.Id)
		nombreVerif = ctx.Query("nombre") != "" && ctx.Query("nombre") != product.Nombre
		colorVerif = ctx.Query("color") != "" && ctx.Query("color") != product.Color
		precioVerif = ctx.Query("precio") != "" && ctx.Query("precio") != fmt.Sprintf("%f", product.Precio)
		stockVerif = ctx.Query("stock") != "" && ctx.Query("stock") != fmt.Sprintf("%d", product.Stock)
		codigoVerif = ctx.Query("codigo") != "" && ctx.Query("codigo") != product.Codigo
		publicadoVerif = ctx.Query("publicado") != "" && ctx.Query("publicado") != fmt.Sprintf("%t", product.Publicado)
		fechaCreacionVerif = ctx.Query("fechaCreacion") != "" && ctx.Query("fechaCreacion") != product.FechaCreacion

		if idVerif || nombreVerif || colorVerif || precioVerif || stockVerif || codigoVerif || publicadoVerif || fechaCreacionVerif {
			continue
		}

		filtrados = append(filtrados, product)

	}

	ctx.JSON(200, filtrados)
}

// Clase 1 Ejercicio 2 Parte 2
func GetById(ctx *gin.Context) {
	producto, err := services.GetById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, producto)
}

// Clase 2 Ejercicio 1 Parte 1
func validateToken(ctx *gin.Context) bool {
	if token := ctx.GetHeader("Authorization"); token != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
		return false
	}
	return true
}

// Clase 2 Ejercicio 1 Parte 1
func Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var producto models.Product

		if !validateToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&producto); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.Field())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.Field())
					}
				}
				ctx.JSON(404, result)
				return
			}
		}

		models.LastId++
		producto.Id = models.LastId
		models.Products = append(models.Products, producto)
		ctx.JSON(http.StatusOK, producto)
	}
}

func main() {
	router := gin.Default()

	// Clase 1 Ejercicio 1 Parte 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	productos := router.Group("/productos")
	{
		// Clase 1 Ejercicio 2 Parte 1
		productos.GET("/", GetAll)
		// Clase 1 Ejercicio 2 Parte 2
		productos.GET("/:id", GetById)
		// Clase 2 Ejercicio 1 Parte 1
		productos.POST("/", Create())
	}

	router.Run(":8080")
}
