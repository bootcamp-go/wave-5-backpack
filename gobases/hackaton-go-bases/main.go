package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file := file.File{Path: "./tickets.csv"}

	tickets, err := file.Read()
	if err != nil {
		log.Println(err)
	}

	// Funcion para obtener tickets del archivo csv
	book := service.NewBookings(tickets)

	nuevoTicket := service.Ticket{
		Names:       "cristian",
		Email:       "cristian@email.com",
		Destination: "misiones",
		Date:        "19:00",
		Price:       700,
	}

	ticket, err := book.Create(nuevoTicket)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("ticket creado: %v\n", ticket)

	read, err := book.Read(1001)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("leo %+v\n", read)

	updated, err := book.Update(1001, service.Ticket{Names: "Nahuel", Email: "otro@email.com"})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("ticket actualizado: %+v\n", updated)

	deleted, err := book.Delete(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("id deleted: %v\n", deleted)

	file.Write(book.ReadAll())
}
