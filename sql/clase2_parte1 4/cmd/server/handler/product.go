package handler

import (
	"errors"
	"net/http"
	"strconv"

	"clase2_parte1/internal/domain"
	"clase2_parte1/internal/products"
	"clase2_parte1/pkg/web"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound               = errors.New("product not found")
	ErrAlreadyExists          = errors.New("product already exists")
	ErrProductRecordsNotFound = errors.New("product records not found for the provided product id")
)

type Product struct {
	productService products.Service
}

func NewProduct(productService products.Service) *Product {
	return &Product{
		productService: productService,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll()
		if err != nil {
			return
		}
		web.Success(c, http.StatusOK, products)
	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve 'id' param from URL")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: provided id '%s' is not an integer", id)
			return
		}
		product, err := p.productService.GetOne(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productToSave domain.Product
		// Asociar el contenido del body a los campos de la estructura Product
		if err := c.ShouldBindJSON(&productToSave); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		savedProduct, err := p.productService.Store(productToSave)
		if err != nil {
			if err == ErrAlreadyExists {
				web.Error(c, http.StatusConflict, "error: %s", err.Error())
				return
			}
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusCreated, savedProduct)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve 'id' param from URL")
			return
		}
		_, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: provided id '%s' is not an integer", id)
			return
		}

		var productToUpdate domain.Product
		if err := c.ShouldBindJSON(&productToUpdate); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}
		updatedProduct, err := p.productService.Update(productToUpdate)
		if err != nil {
			if err.Error() == "product not found" {
				web.Error(c, http.StatusNotFound, "error: %s", err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, updatedProduct)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve 'id' param from URL")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: provided id '%s' is not an integer", id)
			return
		}
		err = p.productService.Delete(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}

func (p *Product) GetFullData() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve 'id' param from URL")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: provided id '%s' is not an integer", id)
			return
		}
		productsAndWarehouses, err := p.productService.GetFullData(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, productsAndWarehouses)
	}
}

func (p *Product) GetOneWithcontext() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve 'id' param from URL")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: provided id '%s' is not an integer", id)
			return
		}
		product, err := p.productService.GetOneWithcontext(c, int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)
	}
}
