package handler

import (
	"clase4_repaso/internal/products"
	"fmt"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": "something went wrong"})
			return
		}
		c.JSON(200, p)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "failed binding json"})
			return
		}

		if req.Name == "" {
			c.JSON(400, gin.H{"error": "el nombre del producto es requerido"})
			return
		}

		if req.Type == "" {
			c.JSON(400, gin.H{"error": "el tipo de producto es requerido"})
			return
		}

		if req.Count == 0 {
			c.JSON(400, gin.H{"error": "la cantidad es requerido"})
			return
		}

		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "el precio es requerido"})
			return
		}

		p, err := p.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			c.JSON(409, gin.H{"error": fmt.Sprintf("fail storing product: %s", err.Error())})
			return
		}

		c.JSON(201, p)
	}
}
