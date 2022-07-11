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

type Product struct{
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

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

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

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

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

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

		var req request
	
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