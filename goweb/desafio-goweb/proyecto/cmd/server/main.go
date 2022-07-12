package main

import (
	"log"
	"proyecto/cmd/server/handlers"
	"proyecto/internal/ticket"
	"proyecto/pkg/readFile"

	"github.com/gin-gonic/gin"
)

func main() {

	bd, err := readFile.LoadTicketsFromFile("tickets.csv")
	if err != nil {
		log.Fatal("No se pudo levantar la base de datos")
	}
	repo := ticket.NewRepository(bd)
	service := ticket.NewService(repo)

	h := handlers.NewTicketHandler(service)
	server := gin.Default()

	gTicket := server.Group("ticket")
	{
		gTicket.GET("/getByCountry/:dest", h.GetTicketsByCountry())
		gTicket.GET("/getAverage/:dest", h.AverageDestination())
	}

	server.Run()

}
