package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/go-testing/internal/products"
	"github.com/bootcamp-go/go-testing/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service products.Service
}

type request struct {
	Nombre string  `json:"nombre"`
	Stock  int     `json:"stock"`
	Precio float64 `json:"precio"`
}

func NewProducts(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, ""))
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, ""))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, products, ""))
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, ""))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		product, err := p.service.Update(int(id), req.Nombre, req.Stock, req.Precio)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusNoContent, web.NewResponse(204, nil, ""))
	}
}
