package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nombre string = "abelardo"

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var sliceProductos = []Producto{
	{
		Id:            1,
		Nombre:        "MAC",
		Color:         "Plateada",
		Precio:        3000,
		Stock:         2,
		Codigo:        "ABCE1235",
		Publicado:     true,
		FechaCreacion: "06/07/2022",
	},
	{
		Id:            2,
		Nombre:        "Lenovo",
		Color:         "Plateada",
		Precio:        1000,
		Stock:         1,
		Codigo:        "ABCD1234",
		Publicado:     true,
		FechaCreacion: "06/07/2022",
	},
	{
		Id:            3,
		Nombre:        "Sony",
		Color:         "Plateada",
		Precio:        2000,
		Stock:         1,
		Codigo:        "ABCF1236",
		Publicado:     false,
		FechaCreacion: "06/07/2022",
	},
}

var mapaProductos = map[string]string{
	"1": "MAC",
	"2": "Lenovo",
	"3": "Sony",
}

// var jsonProductos, err = json.Marshal(sliceProductos)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, sliceProductos)
}

func BuscarProducto(c *gin.Context) {
	producto, ok := mapaProductos[c.Param("id")]
	if ok {
		c.JSON(http.StatusOK, gin.H{c.Param("id"): producto})
	} else {
		c.JSON(http.StatusNotFound, gin.H{c.Param("id"): ""})
	}
}

func Buscar(c *gin.Context) {

	var filtrados []Producto

	if c.Query("precio") == "" || c.Query("asc") == "" {
		c.JSON(http.StatusOK, gin.H{"Productos": filtrados})
		return
	}

	if value, err := strconv.ParseFloat(c.Query("precio"), 8); err != nil {
		c.String(http.StatusNotFound, "Introdujo una query incorrecta")
	} else {

		for _, p := range sliceProductos {
			if c.Query("asc") == "true" && p.Precio > value {
				filtrados = append(filtrados, p)
			}
			if c.Query("asc") == "false" && p.Precio < value {
				filtrados = append(filtrados, p)
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{"Productos Publicados": filtrados})

}

func main() {

	s := gin.Default()

	s.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": nombre})
	})

	// s.GET("/productos", GetAll)

	urlsProductos := s.Group("/productos")

	urlsProductos.GET("/", GetAll)
	urlsProductos.GET("/:id", BuscarProducto)
	urlsProductos.GET("/buscar", Buscar)

	s.Run()

}
