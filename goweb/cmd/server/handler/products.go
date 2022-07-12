package handler

import (
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
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
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token Invalido"))
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	if errors := checkEmpty(req); len(errors) != 0 {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errors...))
		return
	}

	product, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Codigo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, product, ""))

}

func (p Product) Update(ctx *gin.Context) {

	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token Invalido"))
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	if errors := checkEmpty(req); len(errors) != 0 {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errors...))
		return
	}
	product, err := p.service.Update(id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, product)

}
func (p Product) UpdatePrecioStock(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	var req requestPatch

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	product, err := p.service.UpdatePrecioStock(id, req.Precio, req.Stock)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p Product) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	products, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	// Response
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
}

func (p Product) GetByID(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusOK, nil, err.Error()))
		return
	}

	product, err := p.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
}

func (p Product) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	deleted, err := p.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	res := fmt.Sprintf("el ID: %v fue eliminado", deleted)

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, res, ""))
}

func checkEmpty(req request) []string {
	var errors []string

	if req.Nombre == "" {
		errors = append(errors, "falta el campo 'nombre'")
	}

	if req.Precio == 0 {
		errors = append(errors, "falta el campo 'precio'")
	}

	if req.Stock == 0 {
		errors = append(errors, "falta el campo 'stock'")
	}

	if req.Color == "" {
		errors = append(errors, "falta el campo 'color'")
	}

	if req.Codigo == "" {
		errors = append(errors, "falta el campo 'codigo'")
	}

	if req.FechaCreacion == "" {
		errors = append(errors, "Falta el campo'fechaCreacion")
	}

	return errors
}
