package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/internal/productos"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreación string  `json:"fecha_creacion"`
}

type Product struct {
	service productos.Service
}

func NewProduct(s productos.Service) *Product {
	return &Product{service: s}
}

// List products godoc
// @Summary List Products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Succes 200 {object} web.Response
// @Router /productos [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// List products by Id godoc
// @Summary List Product by Id
// @Tags Products
// @Description get product by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "ProductId to view"
// @Succes 200 {object} web.Response
// @Router /productos/{id} [get]
func (p *Product) GetForId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Token Invalido"))
			return
		}
		p, err := p.service.GetForId(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

//StoreProducts godoc
//@Summary Store Products
//@Tags Products
//@Description store products
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param product body request true "Product to store"
//@Succes 200 {object} web.Response
//@Router /productos [POST]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El producto no puede ser vacio"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El color no puede ser vacio"))
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio no puede ser menor a 0"))
			return
		}
		if req.Stock < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El stock no puede ser menor a 0"))
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo no puede ser vacio"))
			return
		}
		if req.FechaCreación == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha de creacion no puede ser vacio"))
			return
		}
		p, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreación)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description endpoint to update a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ProductId to update"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /productos/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id invalido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El producto no puede ser vacio"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El color no puede ser vacio"))
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio no puede ser menor a 0"))
			return
		}
		if req.Stock < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El stock no puede ser menor a 0"))
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo no puede ser vacio"))
			return
		}
		if req.FechaCreación == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha de creacion no puede ser vacio"))
			return
		}
		p, err := p.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreación)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// ParcialUpdateProducts godoc
// @Summary Update price for the product
// @Tags Products
// @Description endpoint to update price for a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "ProductId to update price"
// @Param precio body request false "price data for update"
// @Success 200 {object} web.Response
// @Router /productos/{id} [patch]
func (p *Product) UpdatePrecio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id invalido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "el precio debe ser indicado"))
			return
		}
		p, err := p.service.UpdatePrecio(int(id), req.Precio)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description endpoint to delete a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "ProductId to delete"
// @Success 200 {object} web.Response
// @Router /productos/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token Invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id invalido"))
			return
		}
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
	}
}
