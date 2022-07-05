package main

import (
	"fmt"
	"hackaton-go-bases/internal/service"

	"hackaton-go-bases/internal/file"
)

func main() {
	file := file.File{}
	tickets, _ := file.Read()

	booking := service.NewBookings(tickets)
	t, _ := booking.Read(2)
	booking.Delete(2)
	_, err := booking.Read(2)
	fmt.Println(t)
	fmt.Println(err)

	newticket, err := booking.Create(service.Ticket{1002, "Juan Perez", "test@gmail.com", "test", "15:00", 46500})
	booking.Update(2, newticket)

}
