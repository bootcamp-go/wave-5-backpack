package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name         *string  `json:"name"`
	Color        *string  `json:"color"`
	Price        *float64 `json:"price"`
	Stock        *uint64  `json:"stock"`
	Code         *string  `json:"code"`
	Published    *bool    `json:"published"`
	Warehouse_id *int     `json:"warehouse_id"`
}

type Product struct {
	service products.Service
}

const (
	FIELD_EMPTY = "el campo %s no puede estar vacio"
)

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(500, web.NewRespose(500, nil, "error consultando los datos"))
			return
		}
		ctx.JSON(200, web.NewRespose(200, products, ""))
	}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(500, web.NewRespose(500, nil, err.Error()))
		}
		producto, err := p.service.GetById(ctx, idInt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespose(404, nil, fmt.Sprintf("product with id %d not found", idInt)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespose(200, producto, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body domain.Product true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pr Request

		if err := ctx.ShouldBindJSON(&pr); err != nil {
			ctx.JSON(400, web.NewRespose(400, nil, err.Error()))
			return
		}
		errs := validate(pr)
		if len(errs) > 0 {
			ctx.JSON(400, gin.H{
				"errors": errs,
			})
			return
		}

		NewProduct := domain.Product{
			Name:      *pr.Name,
			Color:     *pr.Color,
			Price:     *pr.Price,
			Stock:     *pr.Stock,
			Code:      *pr.Code,
			Published: *pr.Published,
		}

		created, err := p.service.Store(ctx, NewProduct)
		if err != nil {
			ctx.JSON(400, web.NewRespose(400, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespose(400, created, ""))
	}
}

func (p *Product) UpdateTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pr Request
		idInt, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewRespose(400, nil, err.Error()))
			return
		}
		if err := ctx.ShouldBindJSON(&pr); err != nil {
			ctx.JSON(400, web.NewRespose(400, nil, err.Error()))
			return
		}
		errs := validate(pr)
		if len(errs) > 0 {
			ctx.JSON(400, gin.H{
				"errors": errs,
			})
			return
		}

		NewProduct := domain.Product{
			Id:        idInt,
			Name:      *pr.Name,
			Color:     *pr.Color,
			Price:     *pr.Price,
			Stock:     *pr.Stock,
			Code:      *pr.Code,
			Published: *pr.Published,
		}

		updated, err := p.service.UpdateTotal(ctx, NewProduct)
		if err != nil {
			ctx.JSON(500, web.NewRespose(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewRespose(200, updated, ""))
	}
}

func (p *Product) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pr Request
		idInt, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(500, web.NewRespose(500, nil, err.Error()))
			return
		}
		if err := ctx.ShouldBindJSON(&pr); err != nil {
			ctx.JSON(400, web.NewRespose(400, nil, err.Error()))
			return
		}
		NewProduct := domain.Product{
			Id:        idInt,
			Name:      *pr.Name,
			Color:     *pr.Color,
			Price:     *pr.Price,
			Stock:     *pr.Stock,
			Code:      *pr.Code,
			Published: *pr.Published,
		}
		updated, err := p.service.UpdatePartial(ctx, NewProduct)
		if err != nil {
			ctx.JSON(500, web.NewRespose(500, nil, err.Error()))
		}
		ctx.JSON(200, web.NewRespose(200, updated, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewRespose(404, nil, err.Error()))
		}
		err = p.service.Delete(ctx, idInt)
		if err != nil {
			ctx.JSON(404, web.NewRespose(404, nil, fmt.Sprintf("product: %d not found", idInt)))
			return
		}
		ctx.JSON(204, web.NewRespose(204, "", ""))
	}
}

func validate(product Request) []string {
	var errors []string
	if product.Name == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "nombre").Error())
	}

	if product.Color == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "color").Error())
	}

	if product.Price == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "precio").Error())
	}

	if product.Stock == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "stock").Error())
	}

	if product.Code == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "stock").Error())
	}

	if product.Published == nil {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "publicado").Error())
	}
	return errors
}
