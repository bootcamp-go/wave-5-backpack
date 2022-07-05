package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	file := file.File{Path: "tickets.csv"}
	tickets, err := file.Read()
	if err != nil {
		panic(err)
	}
	bookings := service.NewBookings(tickets)
	_, err = bookings.Create(service.Ticket{
		Id:          1212,
		Names:       "Pepe Perez",
		Email:       "pepeperez@gmail.com",
		Destination: "Cancun",
		Date:        "12:12",
		Price:       100,
	})
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
	}

	ticketFound, err := bookings.Read(100)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
	} else {
		fmt.Printf("%+v\n", ticketFound)
	}

	_, err = bookings.Update(
		1212,
		service.Ticket{
			Id:          1212,
			Names:       "Pepe Perez mod",
			Email:       "pepeperez@gmail.com",
			Destination: "Cancun",
			Date:        "12:12",
			Price:       100,
		},
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	tokenForDeleteId := 1
	deletedId, err := bookings.Delete(tokenForDeleteId)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("Token id: %d deleted\n", deletedId)
	}
	file.Write(bookings.ReadAll())
}
