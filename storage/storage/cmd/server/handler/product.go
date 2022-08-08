package handler

import (
	"github.com/gin-gonic/gin"
	"storage/internal/domain"
	"storage/internal/product"
	"storage/pkg/web"
)

type productRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

func (r productRequest) validateFields() bool {
	return r.Name != "" && r.Type != "" && r.Count > 0 && r.Price > 0
}

type Product struct {
	service product.Service
}

func NewProduct(s product.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		response, err := p.service.GetByName(name)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", response)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req productRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Response(c, 500, "invalid request", nil)
			return
		}
		if !req.validateFields() {
			web.Response(c, 500, "invalid request", nil)
			return
		}
		product := domain.Product{
			Name:  req.Name,
			Type:  req.Type,
			Count: req.Count,
			Price: req.Price,
		}
		result, err := p.service.Store(product)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", result)
	}
}
