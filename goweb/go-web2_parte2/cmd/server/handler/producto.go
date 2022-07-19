package handler

import (
	"net/http"
	"../../../internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id        int    `json:id`
	Nombre    string `json:nombre`
	Color     string `json:color`
	Precio    int    `json:precio`
	Stock     int    `json:stock`
	Codigo    string `json:codigo`
	Publicado bool   `json:publicado`
	Fecha     string `json:fecha`
}

type Productos struct {
	service productos.Service
}

func NuevoProducto(p productos.Service) *Productos {
	return &Productos{service: s}
}

func (p *Productos) GetAll() gin.HandlerFunc {
	token := ctx.GetHeader("token")

	if token != "123" {
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

func (p *Productos) Store() gin.HandlerFunc {

}
