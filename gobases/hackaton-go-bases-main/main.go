package main

import "github.com/bootcamp-go/hackaton-go-bases/internal/service"

func main() {
	var tickets []service.Ticket

	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
