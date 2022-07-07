package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	f := file.New("./tickets.csv")
	tickets, err := f.Read()
	if err != nil {
		fmt.Println(err)
	}
	// Funcion para obtener tickets del archivo csv
	b := service.NewBookings(tickets)

	// CREATE
	t := service.Ticket{
		Id:          1001,
		Names:       "Cristobal Monsalve",
		Email:       "cmonsalve@gmail.com",
		Destination: "Chile",
		Date:        "17:00",
		Price:       100,
	}
	_, er := b.Create(t)
	if er != nil {
		fmt.Println("Error en create")
	}
	f.WriteAll(b.Tickets)

	// READ
	t2, err := b.Read(1000)
	if err != nil {
		fmt.Println("Error en read")
	}
	fmt.Println(t2)

	//UPDATE
	t3 := service.Ticket{
		Id:          1000,
		Names:       "Juan",
		Email:       "juan@gmail.com",
		Destination: "Argentina",
		Date:        "18:00",
		Price:       120,
	}
	_, err2 := b.Update(1000, t3)
	if err2 != nil {
		fmt.Println("Error en update")
	}
	f.WriteAll(b.Tickets)

	//DELETE
	_, err3 := b.Delete(998)
	if err3 != nil {
		fmt.Println("Error en delete")
	}
	f.WriteAll(b.Tickets)
}
