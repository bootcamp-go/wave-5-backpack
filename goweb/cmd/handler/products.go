package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
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

func (p Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, products)

	}
}

func (p Product) GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := p.service.GetProduct(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
		return
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			var vErrors validator.ValidationErrors
			messageError := "Los siguientes campos son requeridos:"
			if errors.As(err, &vErrors) {
				for _, vE := range vErrors {
					messageError += fmt.Sprintf(" %s", vE.Field())
				}
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": messageError,
			})
			return
		}

		product, err := p.service.Store(r.Nombre, r.Color, r.Precio, r.Stock, r.Codigo, *r.Publicado)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}

}

func (p *Product) UpdateAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			var vErrors validator.ValidationErrors
			messageError := "Los siguientes campos son requeridos:"
			if errors.As(err, &vErrors) {
				for _, vE := range vErrors {
					messageError += fmt.Sprintf(" %s", vE.Field())
				}
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": messageError,
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ID invalido",
			})
			return
		}

		product, err := p.service.UpdateAll(int(id), r.Nombre, r.Color, r.Precio, r.Stock, r.Codigo, *r.Publicado)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)

	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ID invalido",
			})
			return
		}

		if err = p.service.Delete(int(id)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "No fue posible borrar el producto",
			})
			return
		}

		ctx.JSON(http.StatusOK, "El producto fue eliminado correctamente")

	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": messageError,
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ID invalido",
			})
			return
		}

		product, err := p.service.Update(int(id), rP.Nombre, rP.Precio)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)

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
