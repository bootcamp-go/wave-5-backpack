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
	list, err := LoadTicketsFromFile("./tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(list)
	serv := tickets.NewService(repo)
	ticketHandler := handler.NewTicket(serv)


	r := gin.Default()
	
	pr := r.Group("tickets")
	pr.GET("", ticketHandler.GetAll())
	pr.GET("/getByCountry/:dest", ticketHandler.GetByCountry())
	pr.GET("getAverage:dest", ticketHandler.GetByAverage())

	// Rutas a desarollar:
	
	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
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
