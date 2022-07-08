package handler

import (
	"errors"
	"fmt"
	"goweb/2/tt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	s service.Service
}

type request struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		products, err := h.s.ListAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}

func (h *Handler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": notBinded,
			})
			return
		}

		product, err := h.s.Store(req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}
