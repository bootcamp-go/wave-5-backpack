/*-----------------------------------------------------------------------------*

     Assignment:	Hackaton :	Go Web
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	The objective of this practical guide is to strengthen and
		deepen the concepts seen in Go Web. For this, we are going to
		pose an integrative challenge that will allow us to review the
		topics we have we have studied.

	© Mercado Libre - IT Bootcamp 2022

------------------------------------------------------------------------------*/

package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := transport.LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)
	t := handler.NewTicket(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:
	pr := r.Group("/ticket")
	pr.GET("/getByCountry/:dest", t.GetTicketsByCountry())
	pr.GET("/getAverage/:dest", t.AverageDestination())

	if err := r.Run(); err != nil {
		panic(err)
	}

}
