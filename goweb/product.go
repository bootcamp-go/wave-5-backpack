package main

import (
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

func GetAll(ctx *gin.Context) {
	/*var products []product

	 p1 := newProduct(1, "Televisor 50´", "Negro", 100000, 100, "AR45RD1", true, "20-02-22")
	 p2 := newProduct(1, "Celular S4", "Gris", 50000, 300, "AR44RD4", true, "15-02-19")
	 p3 := newProduct(1, "Monitor 32´", "Negro", 80000, 50, "AR51OD8", true, "10-02-21")

	products = append(products, p1, p2, p3)

	 jsonData, err := json.Marshal(products)

	  if err != nil {
	  	log.Fatal(err)
	  } else {
	  	ctx.JSON(200, "/products.json")
	  }*/

	ctx.File("products.json")

}
