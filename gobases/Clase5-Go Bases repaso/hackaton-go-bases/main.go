package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	f := file.NewFile("./tickets.csv")
	fmt.Println(f)
	tickets, err := f.Read()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tickets)
		b := service.NewBookings(tickets)
		b.Delete(999)
		b.Delete(997)

		t := service.NewTicket(1000, "fdsfsd", "fdsfds", "dfsdf", "hora", 23232323)
		b.Update(1000, t)
		t2 := service.NewTicket(1004, "orueba", "prueba", "prueba", "prueba", 23232323)
		b.Create(t2)
		prueba, _ := b.Read(1004)

		fmt.Println(tickets)
		fmt.Println(prueba)
		f.Write(prueba)
	}

	// Funcion para obtener tickets del archivo csv

}
