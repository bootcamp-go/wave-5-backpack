package main

import (
	"fmt"
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file1 := file.File{"tickets.csv"}
	tickets, _ := file1.Read()
	booking := service.NewBookings(tickets)

	tkt1, _ := booking.Create(service.Ticket{Names: "Pablo Melegatto", Email: "asd@asd", Destination: "Miami", Date: "Today", Price: 0})
	//tkt2, _ := booking.Read(1001)
	fmt.Println(tkt1)
	booking.Delete(1001)
	//fmt.Println(booking.Update(1001, tkt2))

}
