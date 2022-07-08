package handler

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int     `json:"ID" binding:"-"`
	Nombre        string  `json:"Nombre" binding:"required"`
	Color         string  `json:"Color" binding:"required"`
	Precio        float64 `json:"Precio" binding:"required"`
	Stock         int     `json:"Stock" binding:"required"`
	Codigo        string  `json:"Codigo" binding:"required"`
	Publicado     bool    `json:"Publicado" binding:"required"`
	FechaCreacion string  `json:"FechaCreacion" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, products)

	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		var r request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := p.service.Store(r.Nombre, r.Color, r.Precio, r.Stock, r.Codigo, r.Publicado, r.FechaCreacion)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, product)
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

// func GetProduct(ctx *gin.Context) {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		for _, product := range products {
// 			if product.ID == int(id) {
// 				ctx.JSON(http.StatusOK, product)
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusNotFound, "Error 404")
// 	}
// }

// func NewProduct() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if "123456" != ctx.GetHeader("token") {
// 			ctx.JSON(401, "No tiene permisos para realizar la peticion solicitada")
// 			return
// 		}

// 		var p product
// 		if err := ctx.ShouldBindJSON(&p); err != nil {
// 			return
// 		}

// 		p.ID = newID()
// 		products = append(products, p)

// 		ctx.JSON(http.StatusOK, products)
// 	}
// }
