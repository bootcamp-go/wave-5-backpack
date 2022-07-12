package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {
	store := store.NewStore("tickets.csv")
	ticketList, err := store.Read()
	if err != nil {
		panic("Couldn't load tickets")
	}
	repo := tickets.NewRepository(ticketList)
	service := tickets.NewService(repo)
	t := handler.NewTicket(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	rTicket := r.Group("/ticket")
	{
		rTicket.GET("/getByCountry/:dest", t.GetTicketsByCountry)
		rTicket.GET("/getAverage/:dest", t.AverageDestination)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
