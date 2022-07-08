package main

import (
	"web-server/cmd/server/handler"
	"web-server/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	// productsJson, _ := os.ReadFile("./products.json")
	// if err := json.Unmarshal(productsJson, &product); err != nil {
	// 	log.Fatal((err))
	// }
	//Creo un servidor web con 2 middlewares por defecto: logger and recovery middleware
	router := gin.Default()

	pr := router.Group("products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/name/:id", p.UpdateName())
	pr.PATCH("/:id", p.UpdatePrice())
	pr.DELETE("/:id", p.Delete())

	//Creo un handler utilizando la funcion router.GET("endpoint",Handler) donde endpoint es la ruta relativa y handler es la funcion que toma *gin.Context como argumento.
	router.Run()
}
