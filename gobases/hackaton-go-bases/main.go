package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"
)

func main() {
	fmt.Println("Sistema de reservas de pasajes")
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
