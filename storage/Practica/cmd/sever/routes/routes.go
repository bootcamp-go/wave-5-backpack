package routes

import (
	"database/sql"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/sever/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/docs"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {

	r.setGroup()
	r.buildProductRoutes()
	r.buildSwaggerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repo := products.NewRepository(r.db)
	service := products.NewService(repo)
	handler := handler.NewProduct(service)
	prdRoutes := r.rg.Group("/products")
	{
		prdRoutes.GET("/", handler.GetAll())
		prdRoutes.GET("/:id", handler.GetById())
		prdRoutes.POST("/", handler.Store())
		prdRoutes.PUT("/:id", handler.UpdateTotal())
		prdRoutes.PATCH("/:id", handler.UpdatePartial())
		prdRoutes.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildSwaggerRoutes() {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
