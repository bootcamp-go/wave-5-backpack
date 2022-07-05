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
		fmt.Println(err)
	}

	bookings := service.NewBookings(tickets)

	/* ******************* CRUD ******************* */

	/* ***** Crear un nuevo ticket ***** */
	fmt.Println("***** NUEVO TICKET *****")
	ticket := service.Ticket{
		Id:          service.AutoIncrement(bookings.CountAllTickets()),
		Names:       "Daniela",
		Email:       "email@gmail.com",
		Destination: "Urabá",
		Date:        "12:00",
		Price:       1000000,
	}
	fmt.Println(ticket)

	bookings.Create(ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket creado correctamente")
		fmt.Println()
	}

	/* ***** Buscar un ticket por ID ***** */
	fmt.Println("***** BUSCAR TICKET *****")
	ticket, err = bookings.Read(1003)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ticket: %v\n", ticket)
	}
	fmt.Println()

	/* ***** Actualizar un ticket ***** */
	fmt.Println("***** ACTUALIZAR TICKET *****")
	ticket.Date = "15:30"
	bookings.Update(1003, ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ticket %d actualizado correctamente \n", ticket.Id)
		fmt.Printf("Ticket actualizado: %v\n", ticket)	
	}
	fmt.Println()

	/* ***** Eliminar un ticket por ID ***** */
	fmt.Println("***** ELIMINAR TICKET *****")
	bookings.Delete(1022)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ticket %d eliminado correctamente \n", ticket.Id)
		
	}

	/* Error al buscar el Ticket de nuevo */
	fmt.Printf("\n")

	fmt.Println("***** BUSCAR TICKET X2*****")
	ticket, err = bookings.Read(1007)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Ticket: %v\n", ticket)
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("========================================================")
	fmt.Println()

	/* ***** Crear un nuevo ticket X2 ***** */

	fmt.Println("***** NUEVO TICKET *****")
	ticket = service.Ticket{
		Id:          service.AutoIncrement(bookings.CountAllTickets()),
		Names:       "Valentina",
		Email:       "email@gmail.com",
		Destination: "Republica dominicana",
		Date:        "09:00",
		Price:       1200000,
	}
	fmt.Println(ticket)

	ticket, err = bookings.Create(ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket creado correctamente")
		fmt.Println()
	}

	/* ***** Escribir en el archivo CSV ***** */
	err = file.Write(bookings.CountAllTickets())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cambios guardados correctamente")
	}
}
