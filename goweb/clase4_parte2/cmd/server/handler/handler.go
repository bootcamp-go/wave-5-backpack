package handler

import (
	"fmt"
	"strconv"

	"clase4_parte2/internal/products"
	"clase4_parte2/pkg/web"

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

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "el nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			c.JSON(400, web.NewResponse(400, nil, "el tipo de producto es requerido"))
			return
		}

		if req.Count == 0 {
			c.JSON(400, web.NewResponse(400, nil, "la cantidad es requerido"))
			return
		}

		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "el precio es requerido"))
			return
		}
		p, err := p.service.Update(id, req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			c.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}

		c.JSON(200, p)
	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "id incorrecto"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "el nombre del producto es requerido"))
			return
		}

		p, err := p.service.UpdateName(id, req.Name)
		if err != nil {
			c.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}

		c.JSON(200, p)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "id incorrecto"))
			return
		}

		err = p.service.Delete(id)
		if err != nil {
			c.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}

		c.JSON(204, web.NewResponse(204, fmt.Sprintf("se elimino el producto %d", id), ""))
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(204, web.NewResponse(204, nil, "no se encontraron productos"))
			return
		}
		c.JSON(200, p)
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "el nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			c.JSON(400, web.NewResponse(400, nil, "el tipo de producto es requerido"))
			return
		}

		if req.Count == 0 {
			c.JSON(400, web.NewResponse(400, nil, "la cantidad es requerido"))
			return
		}

		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "el precio es requerido"))
			return
		}

		p, err := p.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			c.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, p, ""))
	}
}
