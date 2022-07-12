package main

import (
	"desafio_web/cmd/server/handler"
	"desafio_web/internal/tickets"
	"desafio_web/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	db := storage.NewStore("../../tickets.csv")
	repo := tickets.NewRepository(db)
	serv := tickets.NewService(repo)
	handler := handler.NewUser(serv)
	router := gin.Default()
	rout := router.Group("ticket")
	rout.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
	rout.GET("/getAverage/:dest", handler.GetAverageCountry())
	router.Run()
}
