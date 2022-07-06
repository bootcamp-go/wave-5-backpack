package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {	
  jsonData, err := os.ReadFile("./transactions.json")
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": "no se pudo leer el archivo",
    })
    return
  }

  // type donde guardamos los datos
  var t []models.Transaction

  if err := json.Unmarshal(jsonData, &t); err != nil {
  	log.Printf("error marshalling: %v\n", err)
  	ctx.JSON(http.StatusInternalServerError, gin.H{
  		"error" : err,
  	})
  	return
  }

  log.Println(t)

  // Response
  ctx.JSON(http.StatusOK, t)
}

func GetByID(ctx *gin.Context) {
	id, err :=  strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

  jsonData, err := os.ReadFile("./transactions.json")
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "error leyendo el archivo": err,
    })
    return
  }
  
  
  var transactions []models.Transaction
  if err := json.Unmarshal(jsonData, &transactions); err != nil {
  	ctx.JSON(http.StatusInternalServerError, gin.H{
  		"error unmarshalling" : err,
  	})
  	return
  }

  for _ , t := range transactions {
  	if t.ID == id {
  		ctx.JSON(http.StatusOK, t)
  		return
  	}
  }

  ctx.JSON(http.StatusNotFound, gin.H{
  	"message" : "transaction by ID not found",
  })
}


func main() {
	router := gin.Default()

	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", GetByID)

	if err := router.Run(":8080"); err != nil {
		log.Println("error en el server")
	}
}
