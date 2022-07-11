package handler

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type request struct {
	Nombre    string  `json:"nombre"`
	Color     string  `json:"color"`
	Precio    float64 `json:"precio"`
	Stock     int     `json:"stock"`
	Codigo    string  `json:"codigo"`
	Publicado bool    `json:"publicado"`
}

type requestPatch struct {
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p Product) CreateProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized)
	}

}
