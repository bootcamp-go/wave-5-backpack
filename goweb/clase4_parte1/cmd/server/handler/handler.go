package handler

import (
	"os"

	"clase4_parte1/internal/products"
	"clase4_parte1/pkg/web"

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
		if req.Type == "" {
			c.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}
		if req.Count == 0 {
			c.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}
		if req.Price == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}
		p, err := p.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
	}

}

// Funcion para validar campos de request
func validar(req request) string {
	var response string = "falta el/los campos: "
	if req.Name == "" {
		response += "Name, "
	}
	if req.Type == "" {
		response += "Type, "
	}
	if req.Count == 0 {
		response += "Count, "
	}
	if req.Price == 0 {
		response += "Price, "
	}
	return response
}
