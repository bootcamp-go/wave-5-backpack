package handler

import (
	"errors"
	"fmt"
	"net/http"
	"storage/2/tt/internal/product"
	"storage/2/tt/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	s product.Service
}

type request struct {
	Name  string  `json:"name" binding:"required"`
	Type  string  `json:"type" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Count int     `json:"count" binding:"required"`
}

func NewProduct(s product.Service) *Product {
	return &Product{s: s}
}

func (h *Product) Read() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		product, err := h.s.Read(id)
		if err != nil {
			code := http.StatusNotFound
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

func (h *Product) ReadByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		products, err := h.s.ReadByName(name)
		if err != nil {
			code := http.StatusNotFound
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, products, ""))
	}
}

func (h *Product) ReadAll() gin.HandlerFunc {
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

func (h *Product) Create() gin.HandlerFunc {
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

		product, err := h.s.Create(req.Name, req.Type, req.Price, req.Count)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusCreated
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

func (h *Product) Update() gin.HandlerFunc {
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

		product, err := h.s.Update(ctx, id, req.Name, req.Type, req.Price, req.Count)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

func (h *Product) Delete() gin.HandlerFunc {
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
