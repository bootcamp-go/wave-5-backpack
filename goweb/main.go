package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactions struct {
	Id       uint
	Codigo   string
	Moneda   string
	Monto    float64
	Emisor   string
	Receptor string
	Fecha    string
}

var t []transactions

func main() {

	file, err := os.ReadFile("./transactions.json")
	if err != nil {
		panic(err)
	}

	if e := json.Unmarshal(file, &t); e != nil {
		log.Fatal(e)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Simón",
		})
	})

	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", GetOne)
	router.GET("/transactions/filter", Filter)

	router.Run()
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, t)
}

func GetOne(c *gin.Context) {
	for _, trans := range t {
		id, _ := strconv.Atoi(c.Param("id"))
		if trans.Id == uint(id) {
			c.String(http.StatusFound, "Transaccion %d encontrada.\nDetalles: %v", id, trans)
			return
		}
	}

	c.String(http.StatusNotFound, "La transaccion %s no existe.", c.Param("id"))
}

func Filter(c *gin.Context) {
	id := c.Query("id")
	codigo := c.Query("codigo")
	moneda := c.Query("moneda")
	monto := c.Query("monto")
	emisor := c.Query("emisor")
	receptor := c.Query("receptor")
	fecha := c.Query("fecha")
	filteredList := []transactions{}

	for _, e := range t {
		if fmt.Sprint(e.Id) == id && e.Codigo == codigo && e.Moneda == moneda && fmt.Sprint(e.Monto) == monto && e.Emisor == emisor && e.Receptor == receptor && e.Fecha == fecha {
			filteredList = append(filteredList, e)
		}
	}

	if len(filteredList) < 1 {
		c.String(http.StatusNotFound, "La transacción no existe")
		return
	}

	c.IndentedJSON(http.StatusFound, filteredList)
}
