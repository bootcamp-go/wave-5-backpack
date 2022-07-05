package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"
)

func main() {
	fmt.Println("Sistema de reservas de pasajes")

	// Leemos la información del archivo
	var pasajes file.File
	pasajes.SetPath("./tickets.csv")
	dataPasajes, err := pasajes.Read()

	// Verificamos que se haya podido abrir el archivo
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dataPasajes)
		var tickets []service.Ticket
		// Funcion para obtener tickets del archivo csv
		service.NewBookings(tickets)
	}
}
