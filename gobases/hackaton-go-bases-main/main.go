package main

import (
	"hackaton/internal/file"
	"hackaton/internal/service"
)

func main() {
	var tickets []service.Ticket
	var bookings service.Bookings

	//Leer CSV
	archivo := file.File{Path: "tickets.csv"}
	tickets, err := archivo.Read()

	//Agregar todos los vuelos leidos
	bookings.NewBookings(tickets)
	bookings.Imprimir(err, "AGREGAR TODOS")

	//Crear uno nuevo
	nuevo_ticket := service.Ticket{Names: "Matias Vince", Email: "matiasvince9@gmail.com", Date: "09:12", Destination: "New Zealand", Price: 1200}
	bookings.Create(&nuevo_ticket)
	// bookings.Imprimir(err, "CREAR NUEVO")

	//Eliminar uno
	err = bookings.Delete(98)
	// bookings.Imprimir(err, "ELIMINAR")

	//Modificar uno
	err = bookings.Update(96, service.Ticket{Names: "Ignacio Fletchers", Email: "ifletcherri@nyu.edu.nz", Date: "15:69", Destination: "NZ", Price: 1420})
	// bookings.Imprimir(err, "MODIFICAR")

	//Guardar en archivo
	archivo.Write(nuevo_ticket)

	archivo = file.File{Path: "tickets.csv"}
	tickets, err = archivo.Read()

	bookings.NewBookings(tickets)
	bookings.Imprimir(err, "ARCHIVO DSP DE GUARDAR")
}
