package handler

import (
	"bootcamp/wave-5-backpack/storage/internal/domain"
	"bootcamp/wave-5-backpack/storage/internal/products"
	"bootcamp/wave-5-backpack/storage/pkg/web"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound              = errors.New("product not found")
	ErrAlreadyExists         = errors.New("product already exists")
	ErrProductRecordNotFound = errors.New("product records not found")
)

type requestName struct {
	Name string `json:"name" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productToSave domain.Product
		if err := c.ShouldBindJSON(&productToSave); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		saveProduct, err := p.service.Store(productToSave)
		if err != nil {
			if err == ErrAlreadyExists {
				web.Error(c, http.StatusConflict, "error: %s", err.Error())
				return
			}
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
		}

		web.Success(c, http.StatusCreated, saveProduct)

	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.service.GetAll()
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
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve id params")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		product, err := p.service.GetOne(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)
	}
}

// func (p *Product) GetByName() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req requestName
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			if strings.Contains(err.Error(), "'required' tag") {
// 				c.JSON(http.StatusUnprocessableEntity, err.Error())
// 				return
// 			}

// 			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
// 			return
// 		}

// 		product, err := p.service.GetByName(req.Name)
// 		if err != nil {
// 			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
// 			return
// 		}

// 		web.Success(c, http.StatusOK, product)
// 	}
// }

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve id params")
			return
		}
		_, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		var productToUpdate domain.Product
		if err := c.ShouldBindJSON(&productToUpdate); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		updateProduct, err := p.service.Update(productToUpdate)
		if err != nil {
			if err.Error() == "product not found" {
				web.Error(c, http.StatusNotFound, "error: %s", err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve id params")
			return
		}
		web.Success(c, http.StatusCreated, updateProduct)
	}
}

func (p *Product) GetFullData() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve id params")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		product, err := p.service.GetFullData(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)

	}
}

func (p *Product) GetOneWithContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			web.Error(c, http.StatusInternalServerError, "error: unable to retrieve id params")
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		product, err := p.service.GetOneWithcontext(c, int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, product)

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
		err = p.service.Delete(int(idInt))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
