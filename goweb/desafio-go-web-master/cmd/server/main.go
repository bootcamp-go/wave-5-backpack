package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	db, err := store.LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(db)
	s := tickets.NewService(repo)
	t := handler.NewTickets(s)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("/ticket/", t.GetAll())
	r.GET("/ticket/getByCountry/:dest", t.GetTicketByDestination())
	r.GET("/ticket/getByCountryNum/:dest", t.GetTicketDestinationNum())
	r.GET("/ticket/getAverage/:dest", t.AverageDestination())

	// Rutas a desarollar:
	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}

}
