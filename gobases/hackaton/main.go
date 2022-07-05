package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton/internal/file"
	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton/internal/service"
)

func main() {
	path := file.File{Path: "./tickets.csv"}
	values, err := path.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	bookings := service.NewBookings(values)
	value, err := bookings.Read(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(value)

	_, err2 := bookings.Delete(10)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	value2, err3 := bookings.Update(15, service.Ticket{Id: 15, Names: "A", Email: "A", Destination: "A", Date: "18:10", Price: 105})
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	fmt.Println(value2)

	err4 := bookings.Create(service.Ticket{Id: 1001, Names: "A", Email: "A", Destination: "A", Date: "18:10", Price: 105})
	if err4 != nil {
		fmt.Println(err4.Error())
	}
	path.Write(bookings.Details())
}
