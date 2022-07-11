package main

import (
	"goweb/4/tm/cmd/server/handler"
	"goweb/4/tm/internal/repository"
	"goweb/4/tm/internal/service"
	"goweb/4/tm/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar .env", err)
	}

	fs := store.NewStore(os.Getenv("FILE_PATH"))
	pr := repository.NewRepository(fs)
	ps := service.NewService(pr)
	ph := handler.NewHandler(ps)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	products := router.Group("/products")
	{
		products.POST("", ph.Create())

		products.GET("", ph.ReadAll())
		products.GET("/:id", ph.Read())

		products.PUT("/:id", ph.Update())

		products.PATCH("/:id", ph.UpdateNamePrice())

		products.DELETE("/:id", ph.Delete())
	}

	router.Run()
}
