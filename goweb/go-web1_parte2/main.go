package main

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id int
	Nombre string
	Color string
	Precio int
	Stock int
	Codigo string
	Publicado bool
	Fecha string
}

func main()  {
	router := gin.Default()
	router.GET("/filtro", FiltroProductosHandler)

	router.GET("/productos/:id", ProductosHandler)

	router.Run()
}

func ProductosHandler(c *gin.Context) {
	productos := generarProductos()
	var pr Productos

	id, _ := strconv.Atoi(c.Param("id"))
	find := false
	for _, p := range productos {
		if p.Id == int(id) {
			find = true
			pr = *p
			break
		}
	}

	if !find {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, pr)
	}
}

func FiltroProductosHandler(c *gin.Context)  {
	productos := generarProductos()
	var pr Productos
	if c.ShouldBindQuery(&pr) == nil {
		fmt.Println(pr.Id)
		fmt.Println(pr.Nombre)
		fmt.Println(pr.Color)
		fmt.Println(pr.Precio)
		fmt.Println(pr.Stock)
		fmt.Println(pr.Codigo)
		fmt.Println(pr.Publicado)
		fmt.Println(pr.Fecha)
	}
	var filtrado []*Productos
	for _, p := range productos {
		if pr.Id == p.Id {
			filtrado = append(filtrado, p)
		}
	}
	c.JSON(http.StatusOK, filtrado)
}

func generarProductos() []*Productos {
	products := []*Productos{
		{Id: 1, Nombre: "Chocramo", Color: "Cafe", Precio: 3000, Stock: 100, Codigo: "CASD23", Publicado: true, Fecha: "1992"},
		{Id: 2, Nombre: "Gala", Color: "Cafe", Precio: 2000, Stock: 100, Codigo: "FRE341", Publicado: true, Fecha: "19"},
	}
	return products
}
