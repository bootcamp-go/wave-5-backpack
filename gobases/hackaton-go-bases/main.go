package main

import "fmt"
import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file := file.File{"tickets.csv"}
	tickets, err := file.Read()
	if err != nil {
		fmt.Println(err)
	}

	booking := service.NewBookings(tickets)

	ticket, err := booking.Read(204)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket)

	ticket.Email = "nuevo@email.com"
	_, err = booking.Update(204, ticket)
	if err != nil {
		fmt.Println(err)
	}

	ticket, err = booking.Read(204)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket)

	booking.Create(service.Ticket{1001, "Vuelo", "email@algo.com", "Buenos Aires", "17:85", 452})
	ticket, err = booking.Read(1001)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket)

	booking.Delete(1000)
	ticket, err = booking.Read(1000)
	if err != nil {
		fmt.Println(err)
	}

	tickets, _ = booking.ReadAll()
	file.Write(tickets)
}
