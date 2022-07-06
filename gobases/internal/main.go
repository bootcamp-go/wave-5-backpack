package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/gobases/internal/file"
	"github.com/bootcamp-go/wave-5-backpack/gobases/internal/service"
)

func main() {
	file := file.File{Path: "tickets.csv"}

	tickets, err := file.Read()

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(tickets[1])

	// Funcion para obtener tickets del archivo csv
	newBooking := service.NewBookings(tickets)

	ticket := service.Ticket{
		Id:          1001,
		Names:       "Francisco Monay",
		Email:       "francisco.monay@mercadolibre.com",
		Destination: "Brasil",
		Date:        "11:00",
		Price:       100}

	ticket, err1 := newBooking.Create(ticket)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		err2 := file.Write(ticket, false)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	var tick *service.Ticket
	tick, err3 := newBooking.Read(1001)

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(tick.Detalle())
	}

	ticket.Names = "Pablo Soldano"
	ticket, err4 := newBooking.Update(1001, ticket)

	if err4 != nil {
		fmt.Println(err4)
	} else {
		file.Write(ticket, true)
	}

	//id, err5 := newBooking.Delete(1000)

	//if err5 != nil {
	//	fmt.Println(err5)
	//} else {
	//	file.Write(ticket, true)
	//	fmt.Printf("El registro con ID %d fue eliminado", id)
	//}

}
