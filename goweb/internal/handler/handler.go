package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {	
	transactions, err := read("./transactions.json")
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  // Response
  ctx.JSON(http.StatusOK, transactions)
}

func GetByID(ctx *gin.Context) {
	id, err :=  strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	transactions, err := read("./transactions.json")
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

	// Aca esta logica
  for _ , transaction := range transactions {
  	if transaction.ID == id {
  		ctx.JSON(http.StatusOK, transaction)
  		return
  	}
  }

  ctx.JSON(http.StatusNotFound, gin.H{
  	"message" : "transaction by ID not found",
  })
}

func GetFilter(ctx *gin.Context) {

	transactions, err := read("./transactions.json")
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
	}

	var (
		cod = ctx.Query("cod")
		moneda = ctx.Query("moneda")
		monto, _ = strconv.ParseFloat(ctx.Query("monto"), 64)
		emisor = ctx.Query("emisor")
		receptor = ctx.Query("receptor")
		fecha = ctx.Query("fecha")
	)

  var filTransactions []models.Transaction
  for _ , t := range transactions {
    if t.Cod == cod && t.Moneda == moneda && t.Monto == monto && t.Emisor == emisor && t.Receptor == receptor && t.Fecha == fecha {
      filTransactions = append(filTransactions, t)
    }
  }
  
  ctx.JSON(http.StatusOK, filTransactions)
}

func read(path string) ([]models.Transaction, error){
	var transactions []models.Transaction

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &transactions)

	return transactions, nil
}
