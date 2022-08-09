package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type request struct {
	ID            int     `json:"ID" binding:"-"`
	Nombre        string  `json:"Nombre" binding:"required"`
	Color         string  `json:"Color" binding:"required"`
	Precio        float64 `json:"Precio" binding:"required"`
	Stock         int     `json:"Stock" binding:"required"`
	Codigo        string  `json:"Codigo" binding:"required"`
	Publicado     *int8   `json:"Publicado" binding:"required"`
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

func (p *Product) GetProductByName() gin.HandlerFunc {
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

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			var vErrors validator.ValidationErrors
			messageError := "Los siguientes campos son requeridos:"
			if errors.As(err, &vErrors) {
				for _, vE := range vErrors {
					messageError += fmt.Sprintf(" %s", vE.Field())
				}
			}
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				400, nil, messageError,
			))
			return
		}

		var product = domain.Product{
			Name:   r.Nombre,
			Type:   r.Color,
			Price:  r.Precio,
			Count:  r.Stock,
			Code:   r.Codigo,
			Public: *r.Publicado,
		}

		id, err := p.service.Store(product)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible crear el producto",
			))
			return
		}

		r.ID = id

		ctx.JSON(http.StatusOK, web.NewResponse(200, r, ""))
	}

}
