package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	var file file.File
	// Funcion para obtener tickets del archivo csv
	file.ChangePath("tickets.csv")
	tickets, err := file.Read()
	if err != nil {
		fmt.Println(err)
	}
	bokin := service.NewBookings(tickets)
	ticket1 := service.Ticket{
		Id:          1010,
		Names:       "Cristian Velez",
		Email:       "cristian.velez@mercadolibre.com",
		Destination: "bogota, Colombia",
		Date:        "14:30",
		Price:       500,
	}
	//2 crear
	bokin.Create(ticket1)
	//4 UpDate
	/*
		_, err2 := bokin.Update(1, ticket1)
		if err2 != nil {
			fmt.Println("err2", err2)
		}*/
	//5 Delete
	//bokin.Delete(2)
	//3 Read
	/*
		newTicket, err := bokin.Read(1)
		if err != nil {
			fmt.Println(err)
			} else {
				fmt.Println(newTicket)
			}*/
	file.WriteAll(bokin.ReadAll())
}
