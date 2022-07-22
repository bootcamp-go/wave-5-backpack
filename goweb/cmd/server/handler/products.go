package handler

import (
	"fmt"
	"strconv"
	"web-server/internal/products"
	"web-server/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Nombre    string  `json:"nombre" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Precio    float64 `json:"precio" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Codigo    string  `json:"codigo" binding:"required"`
	Publicado bool    `json:"publicado"`
	Fecha     string  `json:"fecha_creacion" binding:"required"`
}

type Products struct {
	service products.Service
}

func NewProduct(p products.Service) *Products {
	return &Products{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Descriptions get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "error:No se pudo traer el pedido"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Descriptions store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Products) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error al traer la peticion"))
			return
		}
		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Products) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "No se encontro el id"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("el campo '%s' es requerido", valError.Field())))
					return
				}
			}
			ctx.JSON(400, web.NewResponse(400, nil, "error: ocurrio un error"))
			return
		}

		p, err := p.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "error:no se encontro el producto"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))

	}
}

func (p *Products) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "error:id invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("el campo '%s' es requerido", valError.Field())))
					return
				}
			}
			ctx.JSON(400, web.NewResponse(400, nil, "error:ocurrio un error"))
			return
		}

		p, err := p.service.UpdateName(int(id), req.Nombre)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "error:no se encontro el producto a modificar"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Products) UpdatePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "error:id invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, gin.H{
						"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
					})
					return
				}
			}
			ctx.JSON(400, web.NewResponse(400, nil, "error:ocurrio un error"))
			return
		}

		p, err := p.service.UpdatePrice(int(id), req.Precio)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "error:no se encontro el producto"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Products) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "error:id invalido"))
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "error:No se encontro el producto a eliminar"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, "El producto %s fue eliminado correctamente", ""))
	}
}
