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
	Id       uint    `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
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
	router.POST("/transactions/new", NewTransaction)

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

func NewTransaction(c *gin.Context) {
	if token := c.GetHeader("token"); token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	var req transactions
	var errors []string

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Codigo == "" {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Codigo"))
	}

	if req.Moneda == "" {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Moneda"))
	}

	if req.Monto == 0 {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Monto"))
	}

	if req.Emisor == "" {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Emisor"))
	}

	if req.Receptor == "" {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Receptor"))
	}

	if req.Fecha == "" {
		errors = append(errors, fmt.Sprintf("El campo %s es requerido", "Fecha"))
	}

	if len(errors) > 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": errors,
		})
		return
	}

	lastID := t[len(t)-1].Id
	fmt.Println(lastID)
	req.Id = lastID + 1
	t = append(t, req)

	c.JSON(http.StatusAccepted, req)
}
