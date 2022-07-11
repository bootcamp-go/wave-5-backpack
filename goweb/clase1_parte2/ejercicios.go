package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id            int     `form:"id"`
	Nombre        string  `form:"nombre"`
	Color         string  `form:"color"`
	Precio        float64 `form:"precio"`
	Stock         int     `form:"stock"`
	Codigo        string  `form:"codigo"`
	Publicado     bool    `form:"publicado"`
	FechaCreación string  `form:"fecha_creacion"`
}

func main() {
	router := gin.Default()
	router.GET("/filterproductos", FilterProducts)
	router.GET("/products/:id", ProductsHandler)
	router.Run()
}

func FilterProducts(c *gin.Context) {
	listProducts := generarProducts()
	var pp Productos
	if c.ShouldBindQuery(&pp) == nil {
		log.Println(pp.Id)
		log.Println(pp.Nombre)
		log.Println(pp.Color)
		log.Println(pp.Precio)
		log.Println(pp.Stock)
		log.Println(pp.Codigo)
		log.Println(pp.Publicado)
		log.Println(pp.FechaCreación)
	}
	var filter []*Productos
	for _, p := range listProducts {
		if pp.Id == p.Id && pp.Nombre == p.Nombre && pp.Color == p.Color && pp.Precio == p.Precio && pp.Stock == p.Stock && pp.Codigo == p.Codigo && pp.Publicado == p.Publicado && pp.FechaCreación == p.FechaCreación {
			filter = append(filter, p)
			return
		}
	}
	c.JSON(http.StatusOK, filter)
}

func generarProducts() []*Productos {
	products := []*Productos{
		{Id: 1, Nombre: "Papaya", Color: "Verde", Precio: 1500, Stock: 21, Codigo: "1s1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 2, Nombre: "Aguacate", Color: "Verde", Precio: 2500, Stock: 11, Codigo: "1a1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 3, Nombre: "Melon", Color: "Verde", Precio: 1200, Stock: 9, Codigo: "1m1", Publicado: true, FechaCreación: "06/06/2022"},
	}
	return products
}

func ProductsHandler(c *gin.Context) {
	listProducts := generarProducts()
	var product Productos

	id, _ := strconv.Atoi(c.Param("id"))
	find := false
	for _, p := range listProducts {
		if int64(p.Id) == int64(id) {
			find = true
			product = *p
			break
		}
	}
	if !find {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
