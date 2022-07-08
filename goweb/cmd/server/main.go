package main

import (
	"encoding/json"
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


type product struct{
	Id int `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Color string `form:"color" json:"color"`
	Price float64 `form:"price" json:"price"`
	Stock int `form:"stock" json:"stock"`
	Code string `form:"code" json:"code"`
	Publisher bool `form:"publshier" json:"publisher"`
	CreatedAt string`form:"created_at" json:"created_at"`
}

type request struct{
	Id int `json:"-"`
	Name string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int `json:"stock" binding:"required"`
	Code string `json:"code" binding:"required"`
	Publisher bool `json:"publisher"`
	CreatedAt string`json:"created_at" binding:"required"`
}

var lastId int
func save() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		headerToken := ctx.GetHeader("token")
		if headerToken != "123456" {
			ctx.JSON(401, gin.H{
				"error": "No tiene permisos para realizar la petición solicitada.",
			})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p := read()
		
	

	lastId = p[len(p) - 1].Id + 1
	newP := product{
		Id: lastId,
		Name: req.Name,
		Color: req.Color,
		Price: req.Price,
		Stock: req.Stock,
		Code: req.Code,
		Publisher: req.Publisher,
		CreatedAt: req.CreatedAt,
	}

	p = append(p, newP)

	toJson, err := toJSON(p)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	err1 := os.WriteFile("./productos.json", toJson, 0644)

	if err1 != nil {
		fmt.Println("Error:", err1.Error())
		return
	}

	fmt.Println("p: ", p)
	fmt.Println("lastId:", lastId)
	

	req.Id = lastId
	ctx.JSON(200, req)
	}

	
}

func read() []product {
	var p []product
	file, err := os.ReadFile("./productos.json")

	if err != nil {
		fmt.Println("Error abriendo el archivo productos.json")
	}

	if err1 := json.Unmarshal([]byte(file), &p); err1 != nil {
		log.Fatal(err1)
	}

	return p
	
}

func toJSON(p []product)  ([]byte, error){
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}


func getAll(c *gin.Context) {

	params := c.Request.URL.Query()

	p := read()

	if params != nil{

		var productsFilter []product

	for _, p := range p {
		var boolString = ""

		if p.Publisher == true{
			boolString = "true"
		} else {
			boolString = "false"
		}
		if c.Query("publisher") == string(boolString) {
			productsFilter = append(productsFilter, p)
		}
	}

	fmt.Println(p)
	c.JSON(200, gin.H{
		"products": productsFilter,
	})
	} else {
	
	fmt.Println(p)
	c.JSON(200, gin.H{
		"products": p,
	})
}
}

func filterProducts(ctx *gin.Context)  {

	products := read()

	var productsFilter []product

	for _, p := range products {
		var boolString = ""

		if p.Publisher == true{
			boolString = "true"
		} else {
			boolString = "false"
		}
		if ctx.Query("publisher") == string(boolString) {
			productsFilter = append(productsFilter, p)
		}
	}

	ctx.JSON(http.StatusOK, productsFilter)

}


func main() {

	repo := products.NewRepository()
	serv := products.NewService(repo)
	productHandler := handler.NewProduct(serv)

	// Crea un router con gin
	router := gin.Default()
	pr := router.Group("/productos")

	//  “/productos”
	pr.GET("/", productHandler.GetAll())
	pr.GET("/filtrar", filterProducts)
	pr.POST("/", productHandler.Create())
	
	router.Run()// Corremos nuestro servidor sobre el puerto 8080

}
