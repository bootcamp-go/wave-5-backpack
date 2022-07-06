package main

import (
	"fmt"
	"hackaton-go-bases/internal/service"

	"hackaton-go-bases/internal/file"
)

func main() {
	var tickets []service.Ticket
	file := &file.File{Path: "./tickets.csv"}

	tickets, err := file.Read()

	if err != nil {
		panic(err)
	}

	booking := service.NewBookings(tickets)

	fmt.Println(booking.Read(10))
	fmt.Println(booking.Read(2000))

	newticket := service.Ticket(service.Ticket{Id: 1002, Names: "Juan Perez", Email: "test@gmail.com", Destination: "test", Date: "15:00", Price: 46500})
	booking.Create(newticket)

	fmt.Println(newticket)
	fmt.Println(booking.Update(1, newticket))

	t, _ := booking.Read(2)
	booking.Delete(2)
	_, err = booking.Read(2)
	fmt.Println("ticket 2:", t)
	fmt.Println(err)

	t2, _ := booking.Read(1)
	fmt.Println("Ticket 1: ", t2)
	//fmt.Println(err)

	if err != nil {
		booking.Update(10001, newticket)
	}

	tickets = append(tickets, t2)

	fmt.Println(t2)

	file.Write(tickets)

}
