package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	var newFile file.File
	var fileUpdate file.File

	// Funcion para obtener tickets del archivo csv

	tickets, err := newFile.Read()
	if err != nil {
		fmt.Println("error: no se ha cargado el archivo csv")
	}
	newBooking := service.NewBookings(tickets)

	// Ejecución para crear ticket y añadirlo a la lista de tickets

	nuevoTicket := service.Ticket{
		Names:       "Alan Brito",
		Destination: "Chile",
		Date:        "15:15",
		Email:       "prueba@test.com",
		Price:       150,
	}
	nTicket, _ := newBooking.Create(nuevoTicket)

	fileUpdate.Write(nTicket, tickets)

	// Ejecución para Leer ticket a través de su ID

	fmt.Println(newBooking.Read(1001))

	// Ejecución de Update

	ticketToUpdate := service.Ticket{
		Names:       "San Goku",
		Destination: "Japón",
		Date:        "03:00",
		Email:       "prueba@test.com",
		Price:       150,
	}

	ticketUpdate, tickets, _ := newBooking.Update(1001, ticketToUpdate)
	fmt.Println(ticketUpdate)
	fileUpdate.Write(service.Ticket{}, tickets)

	// Ejecución de Delete

	idDelete := 999
	dTicket, _ := newBooking.Delete(idDelete)
	fileUpdate.Write(service.Ticket{}, dTicket)

}
