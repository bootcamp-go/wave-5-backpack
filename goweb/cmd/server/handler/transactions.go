package handler

import (
	"errors"
	"goweb/cmd/server/utils"
	"goweb/cmd/server/utils/filters"
	"goweb/internal/domain"
	"goweb/internal/transactions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const SECRET_TOKEN = "1234567"

type request struct {
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
	Reciever string  `json:"reciever" binding:"required"`
}

type Transaction struct {
	ser transactions.Service
}

func NewTransaction(ser transactions.Service) Transaction {
	return Transaction{ser}
}

func (t *Transaction) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := t.ser.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var filteredTransactions = make([]domain.Transaction, 0)
		for _, transaction := range transactions {
			if result := filters.PassFilters(
				filters.ContainsString(transaction.Sender, ctx.Query("sender")),
				filters.ContainsString(transaction.Reciever, ctx.Query("reciever")),
				filters.EqString(transaction.Currency, ctx.Query("currency")),
				filters.EqAmount(transaction.Amount, ctx.Query("amount")),
				filters.MinAmount(transaction.Amount, ctx.Query("minAmount")),
				filters.MaxAmount(transaction.Amount, ctx.Query("maxAmount")),
				filters.SameDay(transaction.TransactionDate, ctx.Query("date")),
			); result {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}

		ctx.JSON(200, gin.H{
			"data": filteredTransactions,
		})
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := t.ser.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": transactions,
		})
	}
}

func (t *Transaction) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		searchId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err": "param id must be integer",
			})
			return
		}
		transactionResult, err := t.ser.GetById(searchId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err": "transaction not found",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": transactionResult,
		})
	}
}

func (t *Transaction) CreateTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != SECRET_TOKEN {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access Denied: Token Unauthorized",
			})
			return
		}
		transactionRequest := request{}
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				mesagesErrors := utils.GenerateMessageValidationError(err.(validator.ValidationErrors))
				ctx.JSON(400, gin.H{
					"ValidationErrors": mesagesErrors,
				})
			} else {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		transaction, err := t.ser.Store(transactionRequest.Currency, transactionRequest.Amount, transactionRequest.Sender, transactionRequest.Reciever)
		if err != nil {
			errAmountNotAllowed := transactions.NotAllowedAmountZeroOrNegative{}
			if errors.Is(err, &errAmountNotAllowed) {
				ctx.JSON(400, gin.H{
					"error": "error: amount is zero or below to 0",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(201, transaction)
	}
}
