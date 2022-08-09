package handler

import (
	"goweb/internal/domain"
	"goweb/internal/products"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string  `json:"name" binding:"required"`
	Type         string  `json:"type" binding:"required"`
	Count        int     `json:"count" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Id_warehouse int     `json:"id_warehouse" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) GetOneProductByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: token inválido"))
		}

		name := ctx.Param("name")
		product, err := p.service.GetByName(name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(200, product, ""))

	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: token inválido"))
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(422, web.NewResponse(422, nil, "error: JSON keys required are not included."))
			return
		}

		product := domain.Product{
			Name:  req.Name,
			Type:  req.Type,
			Count: req.Count,
			Price: req.Price,
		}

		product, err := p.service.Store(product)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(201, product, ""))
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: token inválido"))
		}

		products, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(200, products, ""))

	}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: token inválido"))
		}
		id, _ := strconv.Atoi(ctx.Param("id"))
		product, err := p.service.GetFullDataById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(200, product, ""))

	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: token inválido"))
		}

		id, _ := strconv.Atoi(ctx.Param("id"))

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(422, web.NewResponse(422, nil, "error: JSON keys required are not included."))
			return
		}

		product := domain.Product{
			Id:           id,
			Name:         req.Name,
			Type:         req.Type,
			Count:        req.Count,
			Price:        req.Price,
			Id_warehouse: req.Id_warehouse,
		}

		productUpdated, err := p.service.Update(ctx, product)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(201, productUpdated, ""))

	}
}
