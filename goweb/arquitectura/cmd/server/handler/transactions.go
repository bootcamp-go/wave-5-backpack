package handler

import (
	"arquitectura/internal/domain"
	"arquitectura/internal/transactions"
	"arquitectura/pkg/web"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type simplerequest struct {
	TranCode string  `json:"tranCode"`
	Amount   float64 `json:"amount"`
}

func (r simplerequest) validate() bool {
	return r.TranCode != "" || r.Amount > 0
}

type request struct {
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

func (r request) getMissingField() string {
	var errorFields []string
	if r.TranCode == "" {
		errorFields = append(errorFields, "tranCode")
	}
	if r.Currency == "" {
		errorFields = append(errorFields, "currency")
	}
	if r.Amount == 0 {
		errorFields = append(errorFields, "amount")
	}
	if r.Transmitter == "" {
		errorFields = append(errorFields, "transmitter")
	}
	if r.Reciever == "" {
		errorFields = append(errorFields, "receiver")
	}
	if r.TranDate == "" {
		errorFields = append(errorFields, "tranDate")
	}

	if len(errorFields) != 0 {
		var fields string
		for i, f := range errorFields {
			fields += f
			if i != len(errorFields)-1 {
				fields += ","
			}
		}
		return fields
	}

	return ""
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

func isValidToken(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		return false
	}
	return true
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if e := isValidToken(ctx); !e {
			return
		}
		t, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if e := isValidToken(ctx); !e {
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		if e := isValidToken(c); !e {
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "Id inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, "missing the following fields : "+req.getMissingField()))
			return
		}

		t, err := t.service.Update(int(id), req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, t, ""))
	}
}

func (t *Transaction) UpdateFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if e := isValidToken(c); !e {
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req simplerequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if !req.validate() {
			c.JSON(404, web.NewResponse(404, nil, "los campos tranCode y amount no son validos"))
			return
		}

		var tr domain.Transaction
		if req.TranCode != "" {
			t, err := t.service.UpdateTranCode(int(id), req.TranCode)
			if err != nil {
				c.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}
			tr = t
		}
		if req.Amount > 0 {
			t, err := t.service.UpdateAmount(int(id), req.Amount)
			if err != nil {
				c.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}
			tr = t
		}
		c.JSON(200, web.NewResponse(200, tr, ""))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if e := isValidToken(c); !e {
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, fmt.Sprintf("la transaccion con id %d ha sido eliminada", id), ""))
	}
}
