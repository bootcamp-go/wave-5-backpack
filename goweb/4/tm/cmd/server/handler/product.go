package handler

import (
	"errors"
	"fmt"
	"goweb/4/tm/internal/service"
	"goweb/4/tm/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Read() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

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

func (h *Handler) ReadAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

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

func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

		type request struct {
			Name     string  `json:"name" binding:"required"`
			Price    float64 `json:"price" binding:"required"`
			Quantity int     `json:"quantity" binding:"required"`
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

		product, err := h.s.Create(req.Name, req.Price, req.Quantity)
		if err != nil {
			code := http.StatusInternalServerError
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		code := http.StatusOK
		ctx.JSON(code, web.NewResponse(code, product, ""))
	}
}

func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		type request struct {
			Name     string  `json:"name" binding:"required"`
			Price    float64 `json:"price" binding:"required"`
			Quantity int     `json:"quantity" binding:"required"`
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

func (h *Handler) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			ctx.JSON(code, web.NewResponse(code, nil, "id invalido"))
			return
		}

		type request struct {
			Name  string  `json:"name" binding:"required"`
			Price float64 `json:"price" binding:"required"`
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

func (h *Handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			code := http.StatusUnauthorized
			ctx.JSON(code, web.NewResponse(code, nil, "No esta autorizado a realizar esta acción"))
			return
		}

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
