package handler

import (
	"errors"
	"fmt"
	"os"
	"proyecto_meli/internal/products"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, products)
	}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(500, errors.New("ID invalido"))
			return
		}
		product, err := p.service.GetById(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, product)
	}
}

func (p *Product) FilterList() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var name, color, codigo, fecha string
		var stock int = -1
		var id int
		var price float64
		var publicado bool = true
		idQ := c.Query("id")
		if idQ != "" {
			Tem, err := strconv.Atoi(idQ)
			if err != nil {
				c.JSON(500, gin.H{"error": "ID invalido"})
				return
			}
			id = Tem
		}
		nameQ := c.Query("nombre")
		if nameQ != "" {
			name = nameQ
		}
		colorQ := c.Query("color")
		if colorQ != "" {
			color = colorQ
		}
		priceQ := c.Query("precio")
		if priceQ != "" {
			Tem, err := strconv.ParseFloat(priceQ, 64)
			if err != nil {
				c.JSON(500, gin.H{"error": "Precio invalido"})
				return
			}
			price = Tem
		}
		stockQ := c.Query("stock")
		if stockQ != "" {
			Tem, err := strconv.Atoi(stockQ)
			if err != nil {
				c.JSON(500, gin.H{"error": "Stock invalido"})
				return
			}
			stock = Tem
		}
		codigoQ := c.Query("codigo")
		if codigoQ != "" {
			codigo = codigoQ
		}
		publicadoQ := c.Query("publicado")
		if publicadoQ != "" {
			Tem, err := strconv.ParseBool(idQ)
			if err != nil {
				c.JSON(500, gin.H{"error": "Publicado invalido"})
				return
			}
			publicado = Tem
		}
		fechaQ := c.Query("fecha")
		if fechaQ != "" {
			fecha = fechaQ
		}

		products, err := p.service.FilterList(id, name, color, price, stock, codigo, publicado, fecha)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if len(products) > 0 {
			c.JSON(200, products)
			return
		}
		c.JSON(500, gin.H{"error": "No se encontraron productos"})
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(500, err.Error())
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Color == "" {
			errStruct = append(errStruct, "color")
		}
		if request.Precio <= 0 {
			errStruct = append(errStruct, "precio")
		}
		if request.Stock < 0 {
			errStruct = append(errStruct, "stock")
		}
		if request.Codigo == "" {
			errStruct = append(errStruct, "codigo")
		}
		if request.FechaCreacion == "" {
			errStruct = append(errStruct, "fecha_de_creacion")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Los campos %s son requeridos.", strings.Join(errStruct, ", "))})
				return
			} else {
				c.JSON(500, gin.H{"error": fmt.Sprintf("El campo %s es requerido.", errStruct[0])})
				return
			}
		}
		product, err := p.service.Store(request.Nombre, request.Color, request.Precio, request.Stock, request.Codigo, request.Publicado, request.FechaCreacion)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, product)
	}
}
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var request request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Color == "" {
			errStruct = append(errStruct, "color")
		}
		if request.Precio <= 0 {
			errStruct = append(errStruct, "precio")
		}
		if request.Stock < 0 {
			errStruct = append(errStruct, "stock")
		}
		if request.Codigo == "" {
			errStruct = append(errStruct, "codigo")
		}
		if request.FechaCreacion == "" {
			errStruct = append(errStruct, "fecha_de_creacion")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Los campos %s son requeridos.", strings.Join(errStruct, ", "))})
				return
			} else {
				c.JSON(500, gin.H{"error": fmt.Sprintf("El campo %s es requerido.", errStruct[0])})
				return
			}
		}
		product, err := p.service.Update(id, request.Nombre, request.Color, request.Precio, request.Stock, request.Codigo, request.Publicado, request.FechaCreacion)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, product)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}
		err = p.service.Delete(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}

func (p *Product) Update_Name_Price() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}
		var request request

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Precio <= 0 {
			errStruct = append(errStruct, "precio")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Los campos %s son requeridos.", strings.Join(errStruct, ", "))})
				return
			} else {
				c.JSON(500, gin.H{"error": fmt.Sprintf("El campo %s es requerido.", errStruct[0])})
				return
			}
		}
		product, err := p.service.Update_Name_Price(id, request.Nombre, request.Precio)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, product)
	}
}
