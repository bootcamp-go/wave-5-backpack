package handler

import (
	"clase4_parte1/internal/products"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		fmt.Printf("Token desde nuestro env: %s\n", tokenFromEnv)

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		type updateRequest struct {
			Name string `json:"nombre"`
		}

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req updateRequest
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.UpdateName(id, req.Name)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
