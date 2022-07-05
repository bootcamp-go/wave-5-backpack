package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	var f file.File
	f.Path = "./tickets.csv"
	tickets, _ = f.Read()
	var t = service.NewBookings(tickets)
	st := service.Ticket{Names: "Jose", Email: "j@meli.com", Destination: "vzla", Date: "13:20", Price: 500}
	_, err := t.Create(st)
	if err != nil {
		fmt.Println(err)
	}
	f.Write(t.Get())

	// a, err := t.Read(5)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(a)

	// _, err := t.Update(998, st)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err := t.Delete(998)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(t)
}
