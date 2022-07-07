package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            string `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_de_creacion"`
}

func readFile() ([]Producto, error) {
	var productos []Producto
	dataBit, err := os.ReadFile("products.json")
	if err != nil {
		return productos, err
	} else {
		err := json.Unmarshal(dataBit, &productos)
		if err != nil {
			return productos, err
		}
	}

	return productos, nil
}

func filterById(id string) (Producto, error) {
	productos, _ := readFile()
	for _, p := range productos {
		if p.Id == id {
			return p, nil
		}
	}
	return Producto{}, errors.New("No se encontro el producto.")
}

func filterListById(productos []Producto, id string) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Id == id {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByName(productos []Producto, name string) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Nombre == name {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByColor(productos []Producto, color string) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Color == color {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByPrice(productos []Producto, price int) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Precio == price {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByStock(productos []Producto, stock int) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Stock == stock {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByCode(productos []Producto, code string) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Codigo == code {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByPublish(productos []Producto, publish bool) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.Publicado == publish {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func filterListByDate(productos []Producto, date string) (filtroProductos []Producto) {
	for _, p := range productos {
		if p.FechaCreacion == date {
			filtroProductos = append(filtroProductos, p)
		}
	}
	return
}

func getAllHandler(c *gin.Context) {
	productos, err := readFile()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"productos": productos,
		})
	}
}

func filterByIdHandler(c *gin.Context) {
	id := c.Param("id")
	producto, err := filterById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"producto": producto,
		})
	}
}

func filterByFieldsHandler(c *gin.Context) {
	productos, err := readFile()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		id := c.Query("id")
		if id != "" {
			productos = filterListById(productos, id)
		}
		name := c.Query("nombre")
		if name != "" {
			productos = filterListByName(productos, name)
		}
		color := c.Query("color")
		if color != "" {
			productos = filterListByColor(productos, color)
		}
		price := c.Query("precio")
		if price != "" {
			p, _ := strconv.Atoi(price)
			productos = filterListByPrice(productos, p)
		}
		stock := c.Query("stock")
		if stock != "" {
			s, _ := strconv.Atoi(stock)
			productos = filterListByStock(productos, s)
		}
		codigo := c.Query("codigo")
		if codigo != "" {
			productos = filterListByCode(productos, codigo)
		}
		publicado := c.Query("publicado")
		if publicado != "" {
			p, _ := strconv.ParseBool(publicado)
			productos = filterListByPublish(productos, p)
		}
		fecha := c.Query("fecha_de_creacion")
		if fecha != "" {
			productos = filterListByDate(productos, fecha)
		}
	}
	if len(productos) > 0 {
		c.JSON(200, gin.H{
			"productos": productos,
		})
	} else {
		c.JSON(500, gin.H{
			"error": "No hay productos",
		})
	}
}

func main() {
	router := gin.Default()
	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Hello Cristian!",
		})
	})
	router.GET("/productos", getAllHandler)

	productos := router.Group("/producto")
	{
		productos.GET("/", filterByFieldsHandler)
		productos.GET("/:id", filterByIdHandler)
	}
	router.Run()
}
