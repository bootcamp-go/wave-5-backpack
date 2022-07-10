package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func IdNext(tickets []service.Ticket) int {

	mayor := 0
	for _, ticket := range tickets {
		if ticket.Id > mayor {
			mayor = ticket.Id
		}
	}
	return mayor + 1
}

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	archivo_csv := file.File{Path: "tickets.csv"}

	tickets, err := archivo_csv.Read()

	if err != nil {
		fmt.Println(err)
	}

	bookings := service.NewBookings(tickets)

	// Crear ticket

	ticket := service.Ticket{
		Id:          IdNext(bookings.ReadAll()),
		Names:       "Luz Elena",
		Email:       "luz@gmail.com",
		Destination: "Brasil",
		Date:        "16:00",
		Price:       2000,
	}

	bookings.Create(ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se creo un nuevo ticket")
		fmt.Println(ticket)
	}
	fmt.Println("==================================")

	// Leer ticket

	ticket, err = bookings.Read(1006)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket Leido")
		fmt.Println(ticket)
	}
	fmt.Println("==================================")

	// Actualizar Ticket

	ticket.Names = "Elena"
	//ticket.Email = "elena@gmail.com"
	bookings.Update(1004, ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se actualizo correctamente")
		fmt.Println(ticket)
	}
	fmt.Println("==================================")

	//Eliminar Ticket

	bookings.Delete(1003)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket Eliminado")
		fmt.Println(ticket)
	}
	fmt.Println("==================================")

	// Guardar los cambios en el archivo tickets.csv
	err = archivo_csv.Write(bookings.ReadAll())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cambios Guardados")
		fmt.Println(ticket)
	}
	fmt.Println("==================================")

}
