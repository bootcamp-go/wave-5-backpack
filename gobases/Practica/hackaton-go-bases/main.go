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
	ticketCreated, err := bookings.Create(service.Ticket{
		Id:          1212,
		Names:       "Pepe Perez",
		Email:       "pepeperez@gmail.com",
		Destination: "Cancun",
		Date:        "12:12",
		Price:       100,
	})
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
	} else {
		fmt.Printf("%+v\n", ticketCreated)
	}

	ticketFound, err := bookings.Read(100)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
	} else {
		fmt.Printf("%+v\n", ticketFound)
	}

	ticketUpdated, err := bookings.Update(
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
	} else {
		fmt.Printf("%+v\n", ticketUpdated)
	}
	tokenForDeleteId := 1212
	ticketForDelete, err := bookings.Read(tokenForDeleteId)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
	} else {
		fmt.Printf("%+v\n", ticketForDelete)
	}
	deletedId, err := bookings.Delete(tokenForDeleteId)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("Token id: %d deleted\n", deletedId)
	}
	ticketDeleted, err := bookings.Read(tokenForDeleteId)
	fmt.Printf("%+v %v\n", ticketDeleted, err)
	file.Write(bookings.ReadAll())
}
