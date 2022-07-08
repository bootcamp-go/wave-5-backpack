package handler

import (
	"goweb/internal/products"

	"github.com/gin-gonic/gin"
)

type request struct{
	Name string `json:"name"`
	Color string `json:"color"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Code string `json:"code"`
	Publisher bool `json:"publisher"`
}

type Product struct{
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		var req request

		var boolean bool = false
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Color == "" {
			c.JSON(400, gin.H{"error": "El color del producto es requerido"})
			return
		}

		if req.Code == "" {
			c.JSON(400, gin.H{"error": "El código del producto es requerido"})
			return
		}

		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "El precio del producto es requerido"})
			return
		}

		if req.Stock == 0 {
			c.JSON(400, gin.H{"error": "El stock del producto es requerido"})
			return
		}

		if req.Publisher != boolean {
			boolean = true
		}

		productServ, err := p.service.Create(req.Name, req.Color, req.Price, req.Stock, req.Code, boolean)

		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"product": productServ})
	}
}