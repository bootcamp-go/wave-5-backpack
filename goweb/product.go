package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID            int
	Nombre        string
	Color         string
	Precio        float64
	Stock         int
	Codigo        string
	Publicado     bool
	FechaCreacion string
}

func newProduct(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) product {
	return product{
		ID:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fecha,
	}
}

func Read() error {
	jsonData, err := ioutil.ReadFile("./products.json")

	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &products); err != nil {
		return err
	}

	return nil
}

func GetAll(ctx *gin.Context) {
	ctx.File("products.json")
}

func GetFilter(ctx *gin.Context) {
	color := ctx.Query("color")
	precio, _ := strconv.ParseFloat(ctx.Query("precio"), 64)
	var productsFilt []product

	for _, product := range products {
		if product.Color == color && product.Precio > precio {
			productsFilt = append(productsFilt, product)
		}
	}
	ctx.JSON(http.StatusOK, productsFilt)
}

func GetProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, product := range products {
			if product.ID == int(id) {
				ctx.JSON(http.StatusOK, product)
				return
			}
		}
		ctx.JSON(http.StatusNotFound, "Error 404")
	}
}
