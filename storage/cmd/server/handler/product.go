package handler

import (
	"net/http"

	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int     `json:"ID" binding:"-"`
	Nombre        string  `json:"Nombre" binding:"required"`
	Color         string  `json:"Color" binding:"required"`
	Precio        float64 `json:"Precio" binding:"required"`
	Stock         int     `json:"Stock" binding:"required"`
	Codigo        string  `json:"Codigo" binding:"required"`
	Publicado     *bool   `json:"Publicado" binding:"required"`
	FechaCreacion string  `json:"FechaCreacion" binding:"-"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p Product) GetProductByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")

		product, err := p.service.GetProductByName(name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible encontrar el producto",
			))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}
