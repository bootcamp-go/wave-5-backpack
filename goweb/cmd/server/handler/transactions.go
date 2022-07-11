package handler

import (
	"errors"
	"fmt"
	"goweb/cmd/server/utils"
	"goweb/cmd/server/utils/filters"
	"goweb/internal/domain"
	"goweb/internal/transactions"
	"goweb/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
	Reciever string  `json:"reciever" binding:"required"`
}

type requestCurrencyAndAmount struct {
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}

type Transaction struct {
	ser transactions.Service
}

func NewTransaction(ser transactions.Service) Transaction {
	return Transaction{ser}
}

func getAtoiId(ctx *gin.Context) (int, error) {
	idParam, exist := ctx.Params.Get("id")
	badIdMessage := "send a valid id"
	if !exist {
		return 0, errors.New(badIdMessage)
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "send a valid id"))
		return 0, errors.New(badIdMessage)
	}
	return id, nil
}

func generateFieldErrorMessage(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		messageErrors := utils.GenerateMessageValidationError(err.(validator.ValidationErrors))
		messageString := ""
		for i, err := range messageErrors {
			if i == 0 {
				messageString += err
			}
			messageString += "\n" + err
		}
		err = errors.New(messageString)
	}
	return err.Error()
}

func generateServiceErrorWeb(err error) (int, web.Response) {
	errAmountNotAllowed := transactions.NotAllowedAmountZeroOrNegative{}
	if errors.Is(err, &errAmountNotAllowed) {
		web.NewResponse(http.StatusBadRequest, nil, "error: amount is zero or below to 0")
	}
	return web.NewResponse(http.StatusInternalServerError, nil, err.Error())

}

// SearchTransactions godoc
// @Summary Search Transactions
// @Tags Transactions
// @Description Search Transactions
// @Produce  json
// @Param Authorization header string true "token"
// @Param sender query string false "sender"
// @Param reciever query string false "reciever"
// @Param currency query string false "currency"
// @Param amount query float64 false "amount"
// @Param minAmount query float64 false "minAmount"
// @Param maxAmount query float64 false "maxAmount"
// @Param date query time.time false "date"
// @Success 200 {object} web.Response
// @Router /transactions/query [get]
func (t *Transaction) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := t.ser.GetAll()
		if err != nil {
			ctx.JSON(generateServiceErrorWeb(err))
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

		ctx.JSON(web.NewResponse(http.StatusOK, filteredTransactions, ""))
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get all transactions
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := t.ser.GetAll()
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusOK, transactions, ""))
	}
}

// GetByTransactionId godoc
// @Summary Get a transaction by Id
// @Tags Transactions
// @Description Get a transaction by Id
// @Produce  json
// @Param Authorization header string true "token"
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.Response
// @Router /transactions/{id} [get]
func (t *Transaction) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		searchId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "param id must be integer"))
			return
		}
		transactionResult, err := t.ser.GetById(searchId)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusNotFound, nil, "transaction not found"))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusOK, transactionResult, ""))
	}
}

// DeleteTransaction godoc
// @Summary Delete a transaction by id
// @Tags Transactions
// @Description Delete a transaction by id
// @Produce  json
// @Param Authorization header string true "token"
// @Param id path int true "Transaction ID"
// @Success 202 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Success 500 {object} web.Response
// @Router /transactions/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getAtoiId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if err := t.ser.Delete(id); err != nil {
			ctx.JSON(generateServiceErrorWeb(err))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusAccepted, fmt.Sprintf("transaction with id %d was deleted successfully", id), ""))
	}
}

// UpdateCurrencyAndAmount godoc
// @Summary Update Currency and Amount by Transaction Id
// @Tags Transactions
// @Description Update Currency and Amount by Transaction Id
// @Produce  json
// @Param Authorization header string true "token"
// @Param id path int true "Transaction ID"
// @Param product body requestCurrencyAndAmount true "Fields Transaction to Update"
// @Success 202 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Success 500 {object} web.Response
// @Router /transactions/{id} [PATCH]
func (t *Transaction) UpdateCurrencyAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getAtoiId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		transactionRequest := requestCurrencyAndAmount{}
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, generateFieldErrorMessage(err)))
			return
		}

		transaction, err := t.ser.UpdateCurrencyAndAmount(id, transactionRequest.Currency, transactionRequest.Amount)
		if err != nil {
			ctx.JSON(generateServiceErrorWeb(err))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusOK, transaction, ""))
	}
}

// UpdateTransaction godoc
// @Summary Update Transaction by Id
// @Tags Transactions
// @Description Update Transaction by Id
// @Produce  json
// @Param Authorization header string true "token"
// @Param id path int true "Transaction ID"
// @Param product body request true "Transaction to update"
// @Success 202 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Success 500 {object} web.Response
// @Router /transactions/{id} [PUT]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := getAtoiId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		transactionRequest := request{}
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, generateFieldErrorMessage(err)))
			return
		}

		transaction, err := t.ser.Update(id, transactionRequest.Currency, transactionRequest.Amount, transactionRequest.Sender, transactionRequest.Reciever)
		if err != nil {
			ctx.JSON(generateServiceErrorWeb(err))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusOK, transaction, ""))
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 201 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Success 500 {object} web.Response
// @Router /transactions [post]
func (t *Transaction) CreateTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		transactionRequest := request{}
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, generateFieldErrorMessage(err)))
			return
		}
		transaction, err := t.ser.Store(transactionRequest.Currency, transactionRequest.Amount, transactionRequest.Sender, transactionRequest.Reciever)
		if err != nil {
			ctx.JSON(generateServiceErrorWeb(err))
			return
		}
		ctx.JSON(web.NewResponse(http.StatusCreated, transaction, ""))
	}
}
