package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	db := store.NewStore("./tickets.csv")
	repo := tickets.NewRepository(db)
	serv := tickets.NewService(repo)
	ticketHandler := handler.NewTicket(serv)

	r := gin.Default()

	pr := r.Group("tickets")
	pr.GET("", ticketHandler.GetAll())
	pr.GET("/getByCountry/:dest", ticketHandler.GetByCountry())
	pr.GET("/getAverage/:dest", ticketHandler.GetByAverage())

	if err := r.Run(); err != nil {
		panic(err)
	}

}
