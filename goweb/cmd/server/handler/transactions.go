package handler

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/internal/transactions"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Recipient string  `json:"recipient" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func validateFields(err *error) string {
	errAns := ""
	var ve validator.ValidationErrors
	if errors.As(*err, &ve) {
		for _, fe := range ve {
			errAns += fmt.Sprintf("el campo %s es requerido ", fe.Field())
		}
	}
	return errAns
}

func (t *Transaction) GetAll(ctx *gin.Context) {
	if ctx.GetHeader("token") != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}
	issuer := ctx.Query("issuer")
	date := ctx.Query("date")
	ans := []domain.Transaction{}
	transactions, _ := t.service.GetAll()
	for _, transaction := range transactions {
		filter := true
		if issuer != "" && issuer != transaction.Issuer {
			filter = false
		}
		if date != "" && date != transaction.Date {
			filter = false
		}
		if filter {
			ans = append(ans, transaction)
		}
	}
	if len(ans) == 0 {
		ctx.JSON(404, "Nada fue encontrado")
		return
	}
	ctx.JSON(200, ans)
}

func (t *Transaction) Create(ctx *gin.Context) {
	if ctx.GetHeader("token") != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"Error: ": validateFields(&err)})
		return
	}
	transaction, err := t.service.Create(req.Code, req.Currency, req.Amount, req.Issuer, req.Recipient, req.Date)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, transaction)
}

func (t *Transaction) GetOne(ctx *gin.Context) {
	if ctx.GetHeader("token") != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "error: id invalido")
		return
	}
	transaction, err := t.service.GetOne(id)
	if err != nil {
		ctx.JSON(404, "Transaccion no existente")
		return
	}
	ctx.JSON(200, transaction)
}

func (t *Transaction) Update(ctx *gin.Context) {
	if ctx.GetHeader("token") != os.Getenv("TOKEN") {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "error: id invalido")
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"Error: ": validateFields(&err)})
		return
	}
	transaction, err := t.service.Update(id, req.Code, req.Currency, req.Amount, req.Issuer, req.Recipient, req.Date)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error:": err.Error(),
		})
		return
	}
	ctx.JSON(200, transaction)
}

func (t *Transaction) Delete(ctx *gin.Context) {
	if ctx.GetHeader("token") != "elpepe" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "error: id invalido")
		return
	}

	if err := t.service.Delete(id); err != nil {
		ctx.JSON(404, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(200, "Eliminado exitosamente")
}

func (t *Transaction) Update2(ctx *gin.Context) {
	if ctx.GetHeader("token") != "elpepe" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "error: id invalido")
		return
	}
	var req request
	ctx.ShouldBindJSON(&req)
	if req.Code == "" {
		ctx.JSON(400, "error: code requerido")
		return
	}
	if req.Amount == 0 {
		ctx.JSON(400, "error: amount requerido")
		return
	}
	transaction, err := t.service.Update2(id, req.Code, req.Amount)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err})
		return
	}
	ctx.JSON(200, transaction)
}
