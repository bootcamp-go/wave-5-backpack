package handler

import (
	"fmt"
	"goweb/internal/products"
	"goweb/pkg/web"
	"os"
	"strconv"

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

type parcialRequest struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type Product struct{
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}


// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description endpoint to get all products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/ [get]
func (p *Product) GetAll() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, p)
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description endpoint to store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /productos/ [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		var boolean bool = false
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Color == "" {
			c.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
			return
		}

		if req.Code == "" {
			c.JSON(400, web.NewResponse(400, nil, "El código del producto es requerido"))
			return
		}

		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}

		if req.Stock == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
			return
		}

		if req.Publisher != boolean {
			boolean = true
		}

		productServ, err := p.service.Create(req.Name, req.Color, req.Price, req.Stock, req.Code, boolean)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, productServ, ""))
	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description endpoint to update a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ProductId to update"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /productos/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req request

		var boolean bool = false
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Color == "" {
			c.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
			return
		}

		if req.Code == "" {
			c.JSON(400, web.NewResponse(400, nil, "El código del producto es requerido"))
			return
		}

		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}

		if req.Stock == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
			return
		}

		if req.Publisher != boolean {
			boolean = true
		}

		p, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, boolean)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// ParcialUpdateProducts godoc
// @Summary Parcial Update products
// @Tags Products
// @Description endpoint to parcial update a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "ProductId to update"
// @Param product body parcialRequest true "Product to parcial update"
// @Success 200 {object} web.Response
// @Router /productos/{id} [patch]
func (p *Product) ParcialUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req parcialRequest
	
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" && req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "Se debe enviar al menos 1 dato nombre o precio"))
			return
		}
		if req.Price < 0 {
			c.JSON(400, web.NewResponse(400, nil, "El precio no puede ser negativo"))
			return
		}

		p, err2 := p.service.ParcialUpdate(int(id), req.Name, req.Price)
		if err2 != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description endpoint to delete a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "ProductId to delete"
// @Success 200 {object} web.Response
// @Router /productos/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200,fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
	}
}