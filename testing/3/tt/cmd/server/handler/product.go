package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing/3/tt/internal/service"
	"testing/3/tt/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	s service.Service
}

type request struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

type patchRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s: s}
}

// ReadProduct godoc
// @Summary Read product
// @Tags Products
// @Description read product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products/:id [get]
func (h *Handler) Read() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		product, err := h.s.Read(id)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

// ReadProducts godoc
// @Summary Read products
// @Tags Products
// @Description read products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products [get]
func (h *Handler) ReadAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := h.s.ReadAll()
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, products, ""))
	}
}

// CreateProduct godoc
// @Summary Create product
// @Tags Products
// @Description create product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to create"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products [post]
func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			var notBinded string
			if errors.As(err, &ve) {
				notBinded += "Los siguientes campos son requeridos: "
				for _, fe := range ve {
					notBinded += fmt.Sprintf("%s ", fe.Field())
				}
			}

			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, notBinded))
			return
		}

		product, err := h.s.Create(req.Name, req.Price, req.Quantity)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusCreated
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Products
// @Description update product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "product id"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products/:id [put]
func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			var notBinded string
			if errors.As(err, &ve) {
				notBinded += "Los siguientes campos son requeridos: "
				for _, fe := range ve {
					notBinded += fmt.Sprintf("%s ", fe.Field())
				}
			}

			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, notBinded))
			return
		}

		product, err := h.s.Update(id, req.Name, req.Price, req.Quantity)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

// UpdateNamePrice godoc
// @Summary Update name and price
// @Tags Products
// @Description update products partially
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "product id"
// @Param product body patchRequest true "Product to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products/:id [patch]
func (h *Handler) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		var req patchRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			var notBinded string
			if errors.As(err, &ve) {
				notBinded += "Los siguientes campos son requeridos: "
				for _, fe := range ve {
					notBinded += fmt.Sprintf("%s ", fe.Field())
				}
			}

			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, notBinded))
			return
		}

		product, err := h.s.UpdateNamePrice(id, req.Name, req.Price)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description delete product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /products/:id [delete]
func (h *Handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		err = h.s.Delete(id)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, fmt.Sprintf("producto de id %d borrado", id), ""))
	}
}
