package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Monto    float64 `json:"monto"`
	Cod      string  `json:"cod_transaction"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

type requestPatch struct {
	Monto float64 `json:"monto"`
	Cod   string  `json:"cod_transaction"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{service: s}
}

func (t Transaction) CreateTransaction(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
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

	transaction, err := t.service.Store(req.Monto, req.Cod, req.Moneda, req.Emisor, req.Receptor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, transaction, ""))
}

func (t Transaction) Update(ctx *gin.Context) {
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

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	if errors := checkEmpty(req); len(errors) != 0 {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errors...))
		return
	}

	transaction, err := t.service.Update(id, req.Monto, req.Cod, req.Moneda, req.Emisor, req.Receptor)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (t Transaction) UpdateMontoCod(ctx *gin.Context) {
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

	transaction, err := t.service.UpdateMontoCod(id, req.Monto, req.Cod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (t Transaction) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	transactions, err := t.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	// Response
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transactions, ""))
}

func (t Transaction) GetByID(ctx *gin.Context) {
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

	transaction, err := t.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaction, ""))
}

func (t Transaction) Delete(ctx *gin.Context) {
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

	deleted, err := t.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, deleted, ""))
}

// Recibe una request y devuelve un string con los campos faltantes
func checkEmpty(req request) []string {
	var errors []string

	if req.Monto == 0 {
		errors = append(errors, "falta el campo 'monto'")
	}

	if req.Cod == "" {
		errors = append(errors, "falta el campo 'cod_transaction'")
	}

	if req.Moneda == "" {
		errors = append(errors, "falta el campo 'moneda'")
	}

	if req.Emisor == "" {
		errors = append(errors, "falta el campo 'emisor'")
	}

	if req.Receptor == "" {
		errors = append(errors, "falta el campo 'receptor'")
	}

	return errors
}
