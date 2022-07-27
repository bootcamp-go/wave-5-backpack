package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"testing-goweb/internal/domain"
	"testing-goweb/internal/products"
	"testing-goweb/pkg/web"
)

type newRequest struct {
	Precio int `json:"precio"`
	Stock  int `json:"stock"`
}

func (nr newRequest) validar() bool {
	return nr.Precio != 0 || nr.Precio != 0
}

type request struct {
	Nombre        string `json:"nombre" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Precio        int    `json:"precio" binding:"required"`
	Stock         int    `json:"stock" binding:"required"`
	Codigo        string `json:"codigo" binding:"required"`
	Publicado     bool   `json:"publicado" binding:"required"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

func (r request) revisarCampos() string {
	var errorDeCampo []string
	if r.Nombre == "" {
		errorDeCampo = append(errorDeCampo, "nombre")
	}
	if r.Color == "" {
		errorDeCampo = append(errorDeCampo, "color")
	}
	if r.Precio == 0 {
		errorDeCampo = append(errorDeCampo, "precio")
	}
	if r.Stock == 0 {
		errorDeCampo = append(errorDeCampo, "stock")
	}
	if r.Codigo == "" {
		errorDeCampo = append(errorDeCampo, "codigo")
	}
	if r.FechaCreacion == "" {
		errorDeCampo = append(errorDeCampo, "fechaCreacion")
	}
	if len(errorDeCampo) != 0 {
		var campos string
		for i, f := range errorDeCampo {
			campos += f
			if i != len(errorDeCampo)-1 {
				campos += ","
			}
		}
		return campos
	}
	return ""
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		p, err := p.service.GetAll()
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		context.JSON(200, web.NewResponse(200, p, ""))
	}
}
func (p *Product) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req request
		if err := context.Bind(&req); err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		p, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		context.JSON(200, web.NewResponse(200, p, ""))
	}
}
func (p *Product) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi("id")
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, "ID no valido"))
			return
		}
		var req request
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(404, web.NewResponse(404, nil, "faltan los siguientes campos: "+req.revisarCampos()))
			return
		}
		p, err := p.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		context.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Product) UpdatePrecioStock() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, "ID no valido"))
			return
		}
		var nreq newRequest
		if err := context.ShouldBindJSON(&nreq); err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if nreq.validar() {
			context.JSON(404, web.NewResponse(404, nil, "los campos precio y stock no son validos"))
			return
		}
		var pr domain.Product
		if nreq.Precio > 0 && nreq.Stock >= 0 {
			p, err := p.service.UpdatePrecioStock(int(id), nreq.Precio, nreq.Stock)
			if err != nil {
				context.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}
			pr = p
		}
		context.JSON(200, web.NewResponse(200, pr, ""))
	}
}
func (p Product) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, "ID no valido"))
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		context.JSON(200, web.NewResponse(200, fmt.Sprintf("producto con id %d se ha eliminado satisfactoriamente", id), ""))
	}
}
