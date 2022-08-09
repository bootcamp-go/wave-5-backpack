package handler

import (
	"Clase1_Parte1/productos/internal/products"
	"Clase1_Parte1/productos/pkg/web"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `form:"nombre" json:"nombre"`
	Color         string `form:"color" json:"color"`
	Precio        int    `form:"precio" json:"precio"`
	Stock         int    `form:"stock" json:"stock"`
	Codigo        string `form:"codigo" json:"codigo"`
	Publicado     bool   `form:"publicado" json:"publicado"`
	FechaCreacion string `form:"fecha_creacion" json:"fecha_creacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description Get all products
// @Accept json
// @Produce json
// @Param token header string true "Authorization token"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// GetProductByID godoc
// @summary Get product by ID
// @Tags Products
// @Description Get a single product by a given ID
// @Accept json
// @Produce json
// @Param token header string true "Authorization token"
// @Param id path string true "Filter by ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos/{id} [get]
func (p *Product) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		p, err := p.service.GetByID(ctx, id)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// GetProductByName godoc
// @summary Get products by Name
// @Tags Products
// @Description Get a list of products by a given Name
// @Accept json
// @Produce json
// @Param token header string true "Authorization token"
// @Param name path string true "Filter by Name"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos/name/{name} [get]
func (p *Product) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nombre := ctx.Param("name")

		p, err := p.service.GetByName(ctx, nombre)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store product
// @Tags Products
// @Description Store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if err := validateBody(req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		p, err := p.service.Store(ctx, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateProducts godoc
// @Summary Update product
// @Tags Products
// @Description Update product by a given ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Param id path string true "ID to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if err := validateBody(req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		p, err := p.service.Update(ctx, id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateProductsNameAndPrice godoc
// @Summary Update product name and price
// @Tags Products
// @Description Update product name and price by a given ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product name and price to update"
// @Param id path string true "ID to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /productos/{id} [patch]
func (p *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}

		p, err := p.service.UpdateNamePrice(ctx, id, req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete product
// @Tags Products
// @Description Delete product by a given ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path string true "ID to delete"
// @Failure 401 {object} web.Response
// @Router /productos/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		p, err := p.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Product) TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("El token en variable de entorno no está definido")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta el token en la cabecera"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}

		ctx.Next()
	}
}

func validateBody(req request) error {
	if req.Nombre == "" {
		return errors.New("el nombre del producto es requerido")
	}

	if req.Color == "" {
		return errors.New("el color del producto es requerido")
	}

	if req.Precio == 0 {
		return errors.New("el precio del producto es requerido")
	}

	if req.Stock == 0 {
		return errors.New("el stock del producto es requerido")
	}

	if req.Codigo == "" {
		return errors.New("el codigo del producto es requerido")
	}

	if req.FechaCreacion == "" {
		return errors.New("la fecha de creacion del producto es requerido")
	}
	return nil
}
