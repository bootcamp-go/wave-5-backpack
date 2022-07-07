package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func ObtenerSiguienteId(tickets []service.Ticket) int {
	// Supuesto de que estén organizados de la misma manera que lo están en csv
	/*
		if len(b.Tickets) == 0 {
			return 1
		}
		return b.Tickets[len(b.Tickets)-1].Id + 1
	*/

	// Implementación completa
	max := 0
	for _, ticket := range tickets {
		if ticket.Id > max {
			max = ticket.Id
		}
	}
	return max + 1
}

func main() {
	var tickets []service.Ticket

	// Crear un struct File con el path del archivo tickets.csv
	archivo := file.File{Path: "tickets.csv"}

	// Leer todos los tickets del archivo
	tickets, err := archivo.Read()

	if err != nil {
		fmt.Println(err)
	}

	bookings := service.NewBookings(tickets)

	// ===========================================
	// ================== CRUD ===================
	// ===========================================

	// Crear un nuevo ticket

	ticket := service.Ticket{
		Id:          ObtenerSiguienteId(bookings.ReadAll()),
		Names:       "Juan",
		Email:       "prueba@gmail.com",
		Destination: "Cancun",
		Date:        "12:00",
		Price:       100,
	}
	fmt.Println("Creando ticket...")
	fmt.Println(ticket)
	bookings.Create(ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket creado correctamente")
	}
	fmt.Println()

	// Leer un ticket

	ticket, err = bookings.Read(1001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket leido correctamente")
		fmt.Println(ticket)
	}
	fmt.Println()

	// Actualizar un ticket

	ticket.Names = "Juanito"
	ticket.Destination = "Veracruz"
	bookings.Update(ticket.Id, ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket actualizado correctamente")
		fmt.Println(ticket)
	}
	fmt.Println()

	// Eliminar un ticket

	bookings.Delete(ticket.Id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket eliminado correctamente")
	}

	ticket, err = bookings.Read(ticket.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket leido correctamente")
		fmt.Println(ticket)
	}
	fmt.Println()

	// Crear un nuevo ticket

	fmt.Println("Creando ticket...")
	ticket = service.Ticket{
		Id:          ObtenerSiguienteId(bookings.ReadAll()),
		Names:       "Juan",
		Email:       "prueba@gmail.com",
		Destination: "Cancun",
		Date:        "12:00",
		Price:       100,
	}
	fmt.Println(ticket)
	ticket, err = bookings.Create(ticket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket creado correctamente")
	}
	fmt.Println()

	// Guardar los cambios en el archivo tickets.csv
	err = archivo.Write(bookings.ReadAll())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Archivo guardado correctamente")
	}
}
