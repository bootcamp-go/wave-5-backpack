package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Monto    float64 `json:"monto" binding:"required"`
	Cod      string  `json:"cod_transaction" binding:"required"`
	Moneda   string  `json:"moneda" binding:"required"`
	Emisor   string  `json:"emisor" binding:"required"`
	Receptor string  `json:"receptor" binding:"required"`
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := t.service.Store(req.Monto, req.Cod, req.Moneda, req.Emisor, req.Receptor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

func (t Transaction) Update(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := t.service.Update(id, req.Monto, req.Cod, req.Moneda, req.Emisor, req.Receptor)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (t Transaction) UpdateMontoCod(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req requestPatch

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := t.service.UpdateMontoCod(id, req.Monto, req.Cod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (t Transaction) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	transactions, err := t.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusOK, transactions)
}

func (t Transaction) GetByID(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := t.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (t Transaction) GetFilter(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	transactions, err := read("../transactions.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var (
		cod      = ctx.Query("cod")
		moneda   = ctx.Query("moneda")
		monto, _ = strconv.ParseFloat(ctx.Query("monto"), 64)
		emisor   = ctx.Query("emisor")
		receptor = ctx.Query("receptor")
		fecha    = ctx.Query("fecha")
	)

	var filTransactions []models.Transaction
	for _, t := range transactions {
		if t.Cod == cod && t.Moneda == moneda && t.Monto == monto && t.Emisor == emisor && t.Receptor == receptor && t.Fecha == fecha {
			filTransactions = append(filTransactions, t)
		}
	}

	ctx.JSON(http.StatusOK, filTransactions)
}

func read(path string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &transactions)

	return transactions, nil
}

func (t Transaction) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleted, err := t.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": deleted})
}
