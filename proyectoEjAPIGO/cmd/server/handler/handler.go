package handler

import (
	"goweb/internal/transacciones"
	"goweb/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo_transaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda             string  `json:"moneda"`
	Monto              float64 `json:"monto"`
	Emisor             string  `json:"emisor"`
	Receptor           string  `json:"receptor"`
	Fecha_transaccion  string  `json:"fecha_transaccion"`
}

type Transaccion struct {
	service transacciones.Service
}

func NewProduct(s transacciones.Service) *Transaccion {
	return &Transaccion{service: s}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(204, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (t *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}

		if v := validar(req); v != "falta el/los campos: " {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		t, err := t.service.Store(req.Codigo_transaccion, req.Emisor, req.Fecha_transaccion, req.Moneda, req.Receptor, req.Monto)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}

}

func (t *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(404, nil, "Id no encontrado."))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if v := validar(req); v != "falta el/los campos: " {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		t, err := t.service.Update(int(id), req.Codigo_transaccion, req.Moneda, req.Emisor, req.Receptor, req.Fecha_transaccion, req.Monto)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return

		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

func validar(req request) string {
	var response string = "falta el/los campos: "
	if req.Codigo_transaccion == "" {
		response += "Codigo de transaccion, "
	}
	if req.Moneda == "" {
		response += "Moneda, "
	}
	if req.Monto == 0 {
		response += "Monto, "
	}
	if req.Emisor == "" {
		response += "Emisor, "
	}
	if req.Receptor == "" {
		response += "Receptor, "
	}
	if req.Fecha_transaccion == "" {
		response += "Fecha de transaccion, "
	}
	return response
}
