package handler

import (
	"errors"
	"goweb/clase1_clase2/internal/products"
	"goweb/clase1_clase2/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre    string `form:"nombre" json:"nombre"`
	Color     string `form:"color" json:"color"`
	Precio    int    `form:"precio" json:"precio"`
	Stock     int    `form:"stock" json:"stock"`
	Codigo    string `form:"codigo" json:"codigo"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			//ctx.JSON(401, gin.H{"error: ": err.Error()})
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		pr, err := p.service.GetAll(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, pr, ""))
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if err := validateFields(req); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		pr, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, pr, ""))
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if err := validateFields(req); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ps, err := p.service.Update(id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ps, err := p.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

func (p *Product) UpdateFields() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		token := ctx.GetHeader("token")
		if err := validateToken(token); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: el campo nombre es requerido"))
			return
		}

		if req.Precio == 0 {
			ctx.JSON(401, web.NewResponse(401, nil, "error: el campo precio debe ser mayor de 0"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ps, err := p.service.UpdateFields(id, req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, ps, ""))
	}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		producto, err := p.service.GetById(id)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, producto, ""))
	}
}

func validateToken(token string) error {
	if token == "" {
		return errors.New("no ingresó el token y es requerido")
	}
	if token != os.Getenv("TOKEN") {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}

func validateFields(req request) error {
	if req.Nombre == "" {
		return errors.New("el campo nombre es requerido")
	}
	if req.Color == "" {
		return errors.New("el campo color es requerido")
	}
	if req.Precio == 0 {
		return errors.New("el campo precio debe ser mayor de 0")
	}
	if req.Stock == 0 {
		return errors.New("el campo stock debe ser mayor de 0")
	}
	if req.Codigo == "" {
		return errors.New("el campo codigo es requerido")
	}
	if req.Fecha == "" {
		return errors.New("el campo fecha es requerido")
	}
	return nil
}
