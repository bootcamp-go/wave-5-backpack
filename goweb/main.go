package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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


func main() {
	router := gin.Default()

	router.GET("/hola/:name", func(ctx *gin.Context) {
		data := fmt.Sprintf("hola %s", ctx.Param("name"))

		ctx.JSON(200, gin.H{
			"message": data,
		})
	})

	router.GET("/transactions", GetAll)

	if err := router.Run(":8080"); err != nil {
		log.Println("error en el server")
	}
}
