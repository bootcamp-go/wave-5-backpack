package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/storage/internal/domains"
	"github.com/bootcamp-go/storage/internal/products"
	"github.com/bootcamp-go/storage/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID          int     `json:"id"`
	Name        string  `json:"nombre" binding:"required"`
	Type        string  `json:"tipo" binding:"required"`
	Count       int     `json:"cantidad" binding:"required"`
	Price       float64 `json:"precio" binding:"required"`
	WarehouseId int     `json:"warehouse_id" binding:"required"`
}

type requestUpdate struct {
	ID          int     `json:"id"`
	Name        string  `json:"nombre,omitempty"`
	Type        string  `json:"tipo,omitempty"`
	Count       int     `json:"cantidad,omitempty"`
	Price       float64 `json:"precio,omitempty"`
	WarehouseId int     `json:"warehouse_id,omitempty"`
}

type requestName struct {
	Name string `json:"nombre" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (s *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product := domains.Product(req)
		id, err := s.service.Store(c, product)
		if err != nil {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		product.ID = id
		c.JSON(http.StatusCreated, web.NewResponse(product, "", http.StatusCreated))
	}
}

func (s *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		products, err := s.service.GetAll(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(nil, err.Error(), http.StatusInternalServerError))
			return
		}

		if len(products) <= 0 {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, "no hay registros disponibles", http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(products, "", http.StatusOK))
	}
}

func (s *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestName
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product, err := s.service.GetByName(c, req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(product, "", http.StatusOK))
	}
}

func (s *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		var req requestUpdate
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(nil, err.Error(), http.StatusUnprocessableEntity))
			return
		}

		req.ID = id
		err = s.service.Update(c, domains.Product(req))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		product, err := s.service.Get(c, req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(nil, err.Error(), http.StatusInternalServerError))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(product, "", http.StatusOK))
	}
}
