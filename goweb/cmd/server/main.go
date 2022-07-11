package main

import (
	"encoding/json"
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"log"
	"net/http"
	"os"

	"goweb/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/files"
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

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := godotenv.Load(); err != nil{
		fmt.Println("error:", err)
	}

	db := store.NewStore("./productos.json")
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	productHandler := handler.NewProduct(serv)

	docs.SwaggerInfo.Host =os.Getenv("HOST")



	// Crea un router con gin
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := router.Group("/productos")

	//  “/productos”
	pr.Use(TokenAuthMiddleware())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/filtrar", filterProducts)
	pr.POST("/", productHandler.Create())
	pr.PUT("/:id", productHandler.Update())
	pr.PATCH("/:id", productHandler.ParcialUpdate())
	pr.DELETE("/:id", productHandler.Delete())
	
	router.Run()// Corremos nuestro servidor sobre el puerto 8080

}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
 
	if requiredToken == "" {
		log.Fatal("no se encontró el token en variable de entorno")
	}
 
	return func(c *gin.Context) {
		token := c.GetHeader("token")
 
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}
 
		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}
 
		c.Next()
	}
 
 }
 
