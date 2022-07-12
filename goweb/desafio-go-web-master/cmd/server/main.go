package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/domain"
	"desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("././tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	repo := tickets.NewRepository(list)
	s := tickets.NewService(repo)
	h := handler.NewTicket(s)

	// Rutas a desarollar:

	// GET - “/ticket/getByCountry/:dest”
	router.GET("/ticket/getByCountry/:dest", h.GetTicketsByCountry())
	// GET - “/ticket/getAverage/:dest”
	router.GET("/ticket/getAverage/:dest", h.GetTicketsByCountry())

	router.Run(":8080")
	if err := router.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
