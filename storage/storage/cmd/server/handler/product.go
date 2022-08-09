package handler

import (
	"github.com/gin-gonic/gin"
	"storage/internal/domain"
	"storage/internal/product"
	"storage/pkg/web"
	"strconv"
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

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.service.GetAll(c)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", products)
	}
}

func (p *Product) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		response, err := p.service.GetByName(c, name)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		if response.ID == 0 {
			web.Response(c, 200, "", nil)
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
		result, err := p.service.Store(c, product)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", result)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			web.Response(c, 500, "invalid id", nil)
			return
		}
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
			ID:    id,
			Name:  req.Name,
			Type:  req.Type,
			Count: req.Count,
			Price: req.Price,
		}
		result, err := p.service.Update(c, product)
		if err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", result)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			web.Response(c, 500, "invalid id", nil)
			return
		}
		if err = p.service.Delete(c, id); err != nil {
			web.Response(c, 500, err.Error(), nil)
			return
		}
		web.Response(c, 200, "", "eliminado correctamente")
	}
}
