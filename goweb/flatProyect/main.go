package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	LastName   string  `json:"last_name"`
	Mail       string  `json:"mail"`
	Years      int     `json:"years"`
	Tall       float64 `json:"tall"`
	Enable     bool    `json:"enable"`
	CreateDate string  `json:"create_date"`
}

func main() {
	router := gin.Default()
	pr := router.Group("/producto")
	//clase 1_2
	pr.GET("/", GetAll())
	pr.GET("/filtar", Filter())
	//clase 2_1
	pr.POST("/crear", Create())
	pr.POST("/crearautorizado", CreateAuthorize())
	router.Run()
}
func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prods := ReadJson()
		ctx.JSON(http.StatusOK, prods)
	}
}
func Filter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prods := ReadJson()
		var prod Users
		for _, p := range prods {
			if strconv.Itoa(p.Id) == ctx.Query("id") {
				prod = p
				ctx.JSON(http.StatusOK, prod)
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "No se a encontrado el Id solicitado"})
	}
}
func Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prods := ReadJson()
		var req Users
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastId := prods[len(prods)-1].Id
		req.Id = lastId + 1
		WriteJson(req)
		ctx.JSON(http.StatusOK, req)
	}
}
func CreateAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		prods := ReadJson()
		var req Users
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastId := prods[len(prods)-1].Id
		req.Id = lastId + 1
		WriteJson(req)
		ctx.JSON(http.StatusOK, req)
	}
}
func ReadJson() []Users {
	var jsonData []Users
	data, err := os.ReadFile("./productos.js")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
func WriteJson(prod Users) {
	prods := ReadJson()
	prods = append(prods, prod)
	jsonData, err := json.Marshal(prods)
	if err != nil {
		log.Fatal(err)
	}
	e := os.WriteFile("./productos.js", jsonData, 0644)
	if e != nil {
		log.Fatal(err)
	}
}
