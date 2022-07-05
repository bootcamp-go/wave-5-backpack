package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	var fil file.File
	cambio := service.Ticket{
		Id:          1,
		Names:       "Juan Carlos",
		Email:       "hola@hola.com",
		Destination: "CABA",
		Date:        "8demayo",
		Price:       2}

	fil.Path = "./tickets.csv"

	tickets, _ = fil.Read()

	book := service.NewBookings(tickets)

	//fmt.Println(tickets)
	fmt.Println(book)

	j, k := book.Update(9, cambio)
	fmt.Println("_________________________________________________\n")
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(book)

}
