package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	// Leemos el archivo y guardamos los tickets
	var path string = "./tickets.csv"
	file := file.File{Path: path}
	var tickets []service.Ticket
	tickets, _ = file.Read()
	// Guardamos los tickets en la estructura bookings
	bookings := service.NewBookings(tickets)

	// Agregamos un ticket, el ID se agrega automáticamente
	ticket := service.Ticket{Names: "Camilo Calderon", Email: "Camilo@meli.com", Destination: "Colombia", Date: "15:13", Price: 100}
	bookings.Create(ticket)

	// Leemos un ticket que existe y otro que no
	fmt.Println(bookings.Read(1001))
	fmt.Println(bookings.Read(2000))

	// Actualizamos un ticket que existe y otro que no, comprobamos que quede guardado
	ticket.Destination = "Argentina"
	fmt.Println(bookings.Update(1001, ticket))
	fmt.Println(bookings.Read(1001))
	fmt.Println(bookings.Update(2000, ticket))

	// Eliminamos un ticket que existe y otro que no, comprobamos que haya sido eliminado
	fmt.Println(bookings.Delete(1000))
	fmt.Println(bookings.Read(1000))
	fmt.Println(bookings.Delete(2000))

	//Guardamos el nuevo archivo CSV
	file.Write(bookings.Get())
}
