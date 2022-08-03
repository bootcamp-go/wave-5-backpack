package handler

import (
	"net/http"
	"os"

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
