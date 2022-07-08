package main

import (
	"goweb/go-web-II/cmd/handler"
	"goweb/go-web-II/internal/products"

	"github.com/gin-gonic/gin"
)

func main(){
	repository := products.NewRepository() 
	service := products.NewService(repository)
	handler := handler.NewUser(service)
	router := gin.Default()
	rUser := router.Group("users")
	rUser.GET("/", handler.GetAll())
	rUser.POST("/", handler.Store())
	rUser.PUT("/:id", handler.Update())
	rUser.DELETE("/id", handler.Delete())
	router.Run()
	
}

/*
1.- Lo recibe el handler, le pide a service que haga su trabajo
2.- Service implementa la lógica necesaria para el pedido
3.- Sí es que necesitamos consultar la BD, Service se lo va a delegar 
a repository.

4.- Service busca que interceptar si es que hay sucede un error en esa
capa o en repository, para que no llegue al Server
*/ 