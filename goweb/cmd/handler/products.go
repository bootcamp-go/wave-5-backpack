package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get all products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Router /products [get]
func (p Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				400, nil, "No fue posible obtener los productos",
			))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, products, ""))

	}
}

// ListProducts godoc
// @Summary List product
// @Tags Products
// @Description get one product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Failure 404  {object} web.Response
// @Router /products/{id} [get]
func (p Product) GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "ID invalido",
			))
			return
		}

		product, err := p.service.GetProduct(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible encontrar el producto",
			))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
		return
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products, receives a json with the product structure and creates it
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Failure 404  {object} web.Response
// @Router /products [post]
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

		product, err := p.service.Store(r.Nombre, r.Color, r.Precio, r.Stock, r.Codigo, *r.Publicado)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible crear el producto",
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}

}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description update all product, receives id parameter and update it
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Failure 404  {object} web.Response
// @Router /products/{id} [put]
func (p *Product) UpdateAll() gin.HandlerFunc {
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

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "ID invalido",
			))
			return
		}

		product, err := p.service.UpdateAll(int(id), r.Nombre, r.Color, r.Precio, r.Stock, r.Codigo, *r.Publicado)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible actualizar el producto",
			))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))

	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description delete product, receives a id parameter and delete the product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Failure 404  {object} web.Response
// @Router /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "ID invalido",
			))
			return
		}

		if err = p.service.Delete(int(id)); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "No fue posible eliminar el producto"))
			return
		}

		message := fmt.Sprintf("El producto (ID: %d) se elimino correctamente", id)

		ctx.JSON(http.StatusOK, web.NewResponse(200, message, ""))

	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description update product, receive id parameter and update only some fields of the product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "fields to update"
// @Success 200 {object} web.Response
// @Failure 400  {object} web.Response
// @Failure 401  {object} web.Response
// @Failure 404  {object} web.Response
// @Router /products/{id} [patch]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type requestPatch struct {
			Nombre string  `json:"Nombre" binding:"required"`
			Precio float64 `json:"Precio" binding:"required"`
		}

		var rP requestPatch
		if err := ctx.ShouldBindJSON(&rP); err != nil {
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

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "ID invalido",
			))
			return
		}

		product, err := p.service.Update(int(id), rP.Nombre, rP.Precio)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				404, nil, "No fue posible actualizar el producto",
			))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))

	}
}

// func GetFilter(ctx *gin.Context) {
// 	color := ctx.Query("color")
// 	precio, _ := strconv.ParseFloat(ctx.Query("precio"), 64)
// 	var productsFilt []product

// 	for _, product := range products {
// 		if product.Color == color && product.Precio > precio {
// 			productsFilt = append(productsFilt, product)
// 		}
// 	}
// 	ctx.JSON(http.StatusOK, productsFilt)
// }
