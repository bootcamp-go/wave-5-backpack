package main

import (
	"fmt"
	"path/filepath"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func commitChanges(tickets []service.Ticket, repo *file.File) error {
	err := repo.WriteAll(tickets)
	return err
}

func commit(bookings service.Bookings, csvSource *file.File) error {
	tickets, err := bookings.ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		err = commitChanges(tickets, csvSource)
		if err != nil {
			println(err)
			return err
		}
		return nil
	}
}

func main() {
	pathFile, err := filepath.Abs("./tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	fileTicket := file.File{Path: pathFile}
	ticketsRead, err := fileTicket.Read()
	if err != nil {
		fmt.Println(err)
	}
	// Funcion para obtener tickets del archivo csv
	bookings := service.NewBookings(ticketsRead)
	randomRead, err := bookings.Read(23)
	if err != nil {
		fmt.Println(randomRead)
	} else {
		fmt.Println("Lectura realizada exitosamente", randomRead)
	}

	//Ticket Creado
	createdTicket, err := bookings.Create(service.Ticket{Names: "Claudio", Email: "claudio.figueroa@mercadolibre.com", Destination: "Santiago", Date: "13:55", Price: 12500})
	if err != nil {
		fmt.Println(err)
	} else {
		fileTicket.Write(createdTicket)
	}
	// Leer Ticket
	ticketRead, err := bookings.Read(createdTicket.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		// Actualizar Ticket
		ticketRead.Names = "Leopoldo Rodriguez"
		_, err = bookings.Update(ticketRead.Id, ticketRead)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Registro", ticketRead.Id, "modificado correctamente")
		}

	}

	// Borrar Ticket
	deleteTicket, _ := bookings.Create(service.Ticket{Names: "Claudio Figueroa",
		Email:       "claudio.figueroa@mercadolibre.com",
		Destination: "Santiago",
		Date:        "13:55",
		Price:       12500})

	id, err := bookings.Delete(deleteTicket.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Registro", id, "eliminado correctamente")
	}

	// Escribir Cambios a CSV
	err = commit(bookings, &fileTicket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cambios Agregados Correctamente")
	}
}
