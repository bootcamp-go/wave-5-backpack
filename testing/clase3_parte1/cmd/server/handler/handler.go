package handler

import (
	"clase3_parte1/internal/products"

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

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
