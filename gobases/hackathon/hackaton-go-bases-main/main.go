package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {

	path := file.File{
		Path: "tickets.csv",
	}
	values, err := path.Read()
	fmt.Println(values)

	if err != nil {
		fmt.Println(err.Error())
	}

	bookings := service.NewBookings(values)
	valor, err := bookings.Read(6)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(valor)
	t1 := service.Ticket{
		Id:          2,
		Names:       "Yvonne",
		Email:       "yvopintos@gmail.com",
		Destination: "Berlin",
		Date:        "14:24",
		Price:       234,
	}

	bookings.Update(2, t1)
	ticketsTotal := bookings.Detalle()
	err4 := path.Write(ticketsTotal)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println("todo ok")

	// valor2, err5 := bookings.Read(3)
	// if err5 != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(valor2)
	//fmt.Println(bookings.Delete(1000))

}
