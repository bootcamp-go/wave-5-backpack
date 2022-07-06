package main

import (
	"fmt"

	"github.com/anesquivel/hackathon/internal/file"
	"github.com/anesquivel/hackathon/internal/service"
)

func main() {

	// Leyendo file de tickets
	tickets := getDataFromFile()

	// Creando booking con los datos del file
	booking := service.NewBookings(tickets)
	createTicket(booking)

	// jalando nuevos datos
	tickets, _ = booking.ReadAll()
	// Editando el primer elemento de la lista
	ticketToUpdate := creatingTicket()
	editTicket(booking, 1, ticketToUpdate)

	// eliminando un elemento de la lista de tickets
	//deleteTicket(booking, 1)
}

/* Función que obtiene todos los registros del file tickets.csv*/
func getDataFromFile() []service.Ticket {
	defer func() {
		errRecovered := recover()
		if errRecovered != nil {
			fmt.Println("Ocurrió un error:", errRecovered)
		}
	}()
	ticketsFile := file.File{Path: "./tickets.csv"}
	tickets, err := ticketsFile.Read()

	if err != nil {
		panic(err)
	}

	return tickets
}

/* función que crea un ticket con data hardcoreada*/
func creatingTicket() service.Ticket {
	return service.Ticket{
		Names:       "Andrea Esquivel",
		Email:       "ing.andreaesquivel@gmail.com",
		Date:        "Mar 5 Jul 11:00",
		Price:       15500,
		Destination: "Paris, France",
	}
}

/* función que guarda el ticket en la lista local y en el file tickets.csv*/
func createTicket(booking service.Bookings) {
	defer func() {
		errRecovered := recover()
		if errRecovered != nil {
			fmt.Println(errRecovered)
		}
	}()
	_, err := booking.Create(creatingTicket())
	if err != nil {
		panic(err)
	}

	allTickets, errGet := booking.ReadAll()

	if errGet != nil {
		panic(errGet)
	}

	writeOnFile(allTickets, "----> Ticket agregado correctamente.😃")

}

/* funcion que llama a las funcs pertinentes para editar un registro segun su ID*/
func editTicket(booking service.Bookings, id int64, contentToUpdate service.Ticket) {
	defer func() {
		errEdit := recover()
		if errEdit != nil {
			fmt.Println(errEdit)
		}
	}()

	_, err := booking.Update(id, contentToUpdate)
	if err != nil {
		panic(err)
	}
	ticketsUpdated, _ := booking.ReadAll()
	writeOnFile(ticketsUpdated, "----> Ticket actualizado correctamente.✅")
}

/* funcion que sobreescribe el archivo newtickets.csv de acuerdo a la data pasada por
params e imprime el mensaje correspondiente de éxito*/
func writeOnFile(allTickets []service.Ticket, msgSuccess string) {
	defer func() {
		errRecovered := recover()
		if errRecovered != nil {
			fmt.Println(errRecovered)
		}
	}()
	ticketsFile := file.File{Path: "./newtickets.csv"}

	errWriting := ticketsFile.Write(allTickets)
	if errWriting != nil {
		panic(errWriting)
	} else {
		fmt.Println(msgSuccess)
	}

}

/* funcion que elimina un ticket a traves de su id*/
func deleteTicket(booking service.Bookings, id int64) {
	defer func() {
		errDel := recover()
		if errDel != nil {
			fmt.Println(errDel)
		}
	}()
	deleted, err := booking.Delete(id)
	if err != nil || deleted == 1 {
		panic(err)
	}

	ticketsUpdated, errRead := booking.ReadAll()
	if errRead != nil {
		panic(errRead)
	}
	writeOnFile(ticketsUpdated, "----> Ticket eliminado.🗑")

}
