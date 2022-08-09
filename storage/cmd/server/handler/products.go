package handler

import (
	"fmt"
	"net/http"
	"storage/internal/domain"
	"storage/internal/products"
	"storage/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id            int     `json:"-"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"código"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creación"`
}

type Product struct {
	service products.Service
}

func InitProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// ListOneProduct godoc
// @Summary List product
// @Tags Products
// @Description get one product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/:id [get]
func (c *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.NewResponse(400, c, "id no valido")
			return
		}
		p, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// CreateProducts godoc
// @Summary Create products
// @Tags Products
// @Description create products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/create [post]
func (c *Product) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productArray domain.Products
		if err := ctx.ShouldBindJSON(&productArray); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if v := validador(productArray); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		p, err := c.service.CreateProduct(productArray)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusOK, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productArray domain.Products
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if err := ctx.ShouldBindJSON(&productArray); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		p, err := c.service.Update(ctx, productArray, int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}

}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id inválido"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK,
			web.NewResponse(http.StatusOK, nil, fmt.Sprintf("El producto %d ha sido eliminado", id)))
	}

}

func validador(productArray domain.Products) string {
	var response string
	if productArray.Nombre == "" {
		response += "Falta el campo nombre "
	}
	if productArray.Color == "" {
		response += "Falta el campo color "
	}
	if productArray.Precio == 0 {
		response += "Falta el campo precio "
	}
	if productArray.Stock == 0 {
		response += "Falta el campo stock "
	}
	if productArray.Codigo == "" {
		response += "Falta el campo código "
	}
	if productArray.FechaCreacion == "" {
		response += "Falta el campo fecha de creación "
	}
	return response
}
