package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/bookings"
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/tickets"
)

func main() {
	var csvTickets []tickets.Ticket
	csv := file.File{Path: "./tickets.csv"}
	csvTickets, _ = csv.Read()
	book := bookings.NewBookings(csvTickets, csv)
	ticketCreado := tickets.Ticket{
		Id:          12345,
		Names:       "Juan Sebastian",
		Email:       "juansebastian.salazar@mercadolibre.com",
		Destination: "Cali",
		Date:        "20:00",
		Price:       700,
	}
	book.Create(ticketCreado)

	encontrado, _ := book.Read(12345)
	fmt.Println(encontrado)
	updateTicket := tickets.Ticket{
		Id:          54321,
		Names:       "Juan Sebastian",
		Email:       "juansebastian.salazar@mercadolibre.com",
		Destination: "Cali",
		Date:        "20:00",
		Price:       700,
	}

	book.Update(12345, updateTicket)
	postUpdate, _ := book.Read(54321)
	fmt.Println(postUpdate)
	book.Delete(54321)
	postDelete, _ := book.Read(54321)
	fmt.Println(postDelete)

}
