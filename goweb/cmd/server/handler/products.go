package handler

import (
	"fmt"
	"goweb/internal/products"
	"goweb/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Code      string `json:"code"`
	Published bool   `json:"published"`
	Date      string `json:"date"`
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
			c.JSON(401, web.NewResponse(401, nil, "Error: Token invalido"))
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(204, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Error: Token invalido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		// if v := validado(req); v != "" {
		// 	c.JSON(400, web.NewResponse(400, nil, v))
		// 	return
		// }
		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Color == "" {
			c.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}
		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}
		if req.Stock == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El stock es requerido"))
			return
		}
		if req.Code == "" {
			c.JSON(400, web.NewResponse(400, nil, "El codigo es requerido"))
			return
		}
		if req.Date == "" {
			c.JSON(400, web.NewResponse(400, nil, "La fecha  es requerida"))
			return
		}

		p, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.Date)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
	}

}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Error: Token invalido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Color == "" {
			c.JSON(400, gin.H{"error": "El tipo de producto es requerido"})
			return
		}

		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "La cantidad es requerido"})
			return
		}

		if req.Stock == 0 {
			c.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}

		if req.Code == "" {
			c.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}
		if req.Date == "" {
			c.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}

		p, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.Date)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, p)
	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Error: Token invalido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		p, err := p.service.UpdateName(int(id), req.Name)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, p)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Error: Token invalido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
