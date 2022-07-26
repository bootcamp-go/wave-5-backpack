package handler

import (
	"errors"
	"goweb/clase1_clase2/internal/products"
	"goweb/clase1_clase2/pkg/web"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre    string `form:"nombre" json:"nombre"`
	Color     string `form:"color" json:"color"`
	Precio    int    `form:"precio" json:"precio"`
	Stock     int    `form:"stock" json:"stock"`
	Codigo    string `form:"codigo" json:"codigo"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

// GetProducts godoc
// @Summary Get products
// @Tags Products
// @Description Get the full or filtered product list
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param nombre query string false "Filter by name"
// @Param color query string false "Filter by color"
// @Param precio query int false "Filter by price"
// @Param stock query int false "Filter by stock"
// @Param codigo query string false "Filter by code"
// @Param publicado query bool false "Filter by published"
// @Param fecha query string false "Filter by date"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		pr, err := p.service.GetAll(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)

		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, pr, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description Add a new product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if err := validateFields(req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		pr, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, pr, ""))
	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description Update a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Update a product"
// @Param id path string true "Update by ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if err := validateFields(req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ps, err := p.service.Update(id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description Delete a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "Delete by ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ps, err := p.service.Delete(id)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

// UpdateFieldsProducts godoc
// @Summary Update the fields of a product
// @Tags Products
// @Description Update the name and price of a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Update fields of a product"
// @Param id path string true "Update fields by ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products/{id} [patch]
func (p *Product) UpdateFields() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "error: el campo nombre es requerido"))
			return
		}

		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "error: el campo precio debe ser mayor de 0"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ps, err := p.service.UpdateFields(id, req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

// GetProductsById godoc
// @Summary Get products by id
// @Tags Products
// @Description Get a product by id
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "Update fields by ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products/{id} [get]
func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		producto, err := p.service.GetById(id)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, producto, ""))
	}
}

func (p *Product) TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "no ingresó el token y es requerido"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
			return
		}

		ctx.Next()
	}

}

func validateFields(req request) error {
	if req.Nombre == "" {
		return errors.New("el campo nombre es requerido")
	}
	if req.Color == "" {
		return errors.New("el campo color es requerido")
	}
	if req.Precio == 0 {
		return errors.New("el campo precio debe ser mayor de 0")
	}
	if req.Stock == 0 {
		return errors.New("el campo stock debe ser mayor de 0")
	}
	if req.Codigo == "" {
		return errors.New("el campo codigo es requerido")
	}
	if req.Fecha == "" {
		return errors.New("el campo fecha es requerido")
	}
	return nil
}
