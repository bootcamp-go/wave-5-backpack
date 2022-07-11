package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      uint64  `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
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

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
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
		producto, err := p.service.GetById(idInt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespose(404, nil, fmt.Sprintf("product with id %d not found", idInt)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespose(200, producto, ""))
	}
}

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

		created, err := p.service.Store(pr.Name, pr.Color, pr.Price, pr.Stock, pr.Code, pr.Published, pr.Created_at)
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

		updated, err := p.service.UpdateTotal(idInt, pr.Name, pr.Color, pr.Price, pr.Stock, pr.Code, pr.Published, pr.Created_at)
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
		fmt.Println(pr)
		updated, err := p.service.UpdatePartial(idInt, pr.Name, pr.Color, pr.Price, pr.Stock, pr.Code, pr.Published, pr.Created_at)
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
			ctx.JSON(500, web.NewRespose(500, nil, err.Error()))
		}
		producto, err := p.service.Delete(idInt)
		if err != nil {
			ctx.JSON(500, web.NewRespose(500, nil, fmt.Sprintf("product: %d not found", idInt)))
			return
		}
		ctx.JSON(500, web.NewRespose(500, producto, ""))
	}
}

func validate(product Request) []string {
	var errors []string
	if product.Name == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "nombre").Error())
	}

	if product.Color == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "color").Error())
	}

	if product.Price == 0 {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "precio").Error())
	}

	if product.Code == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "stock").Error())
	}
	return errors
}
