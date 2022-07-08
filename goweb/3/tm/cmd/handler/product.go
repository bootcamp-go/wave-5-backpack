package handler

import (
	"errors"
	"fmt"
	"goweb/3/tm/internal/service"
	"net/http"
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
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id invalido",
			})
			return
		}

		product, err := h.s.Read(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (h *Handler) ReadAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		products, err := h.s.ReadAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": notBinded,
			})
			return
		}

		product, err := h.s.Create(req.Name, req.Price, req.Quantity)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id invalido",
			})
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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": notBinded,
			})
			return
		}

		product, err := h.s.Update(id, req.Name, req.Price, req.Quantity)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (h *Handler) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id invalido",
			})
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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": notBinded,
			})
			return
		}

		product, err := h.s.UpdateNamePrice(id, req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (h *Handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id invalido",
			})
			return
		}

		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("producto de id %d borrado", id),
		})
	}
}
