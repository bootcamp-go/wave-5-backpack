package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	//service package
	var tickets []service.Ticket
	//Creando tickets para prubas de Bookings
	tickets = append(tickets, service.Ticket{2002, "Luz Lucumí", "luz.lucumi@hotmail.es", "Chile", "12/02/2022", 2300000})
	tickets = append(tickets, service.Ticket{2003, "Martha", "martha@hotmail.es", "Colombia", "12/22/2022", 2400000})
	tickets = append(tickets, service.Ticket{2004, "Luber", "luber@hotmail.es", "EEUU", "12/10/2022", 2600000})

	//CREATE, READ, UPDATE, DELETE
	//file package
	filePath := &file.File{"./tickets.csv"}
	//leyendo el csv
	filePath.Read()

	//CREATE
	dataWrite := ""
	for _, values := range tickets {
		//Creando tickets
		fmt.Println(values)
		// stringformat = append(stringformat, values)
		dataWrite += string(fmt.Sprint(values.Id, ",", values.Names, ",", values.Email, ",", values.Destination, ",", values.Date, ",", values.Price, "\n"))
		_, err := service.Bookings.Create(service.NewBookings(tickets), values)
		if err != nil {
			fmt.Println(err)
		}
	}

	//Escribo los valores en el csv
	filePath.Write(dataWrite)

	//READ
	dataRead, err := filePath.Read()
	if err != nil {
		fmt.Println(err)
	}
	_, errRead := service.Bookings.Read(service.NewBookings(dataRead), 26)
	if errRead != nil {
		fmt.Println(err)
	}
	fmt.Println(dataRead)

	//UPDATE
	dataUpdate, err := service.Bookings.Update(service.NewBookings(dataRead), 2002, tickets[0])
	dataWriteUpdate := ""
	// stringformat = append(stringformat, values)
	dataWriteUpdate += string(fmt.Sprint(dataUpdate.Id, ",", dataUpdate.Names, ",", dataUpdate.Email, ",", dataUpdate.Destination, ",", dataUpdate.Date, ",", dataUpdate.Price, "\n"))
	if err != nil {
		fmt.Println(err)
	}

	//Escribo los valores en el csv
	filePath.Write(dataWriteUpdate)

	//DELETE
	ticketsDelete, deleteValue, err := service.Bookings.Delete(service.NewBookings(dataRead), 1000)
	if err != nil {
		fmt.Print(err)
	}

	dataWriteDelete := ""
	for _, values := range ticketsDelete {
		//Creando tickets
		fmt.Println(values)
		// stringformat = append(stringformat, values)
		dataWriteDelete += string(fmt.Sprint(values.Id, ",", values.Names, ",", values.Email, ",", values.Destination, ",", values.Date, ",", values.Price, "\n"))
	}

	//Escribo los valores en el csv
	filePath.Write(dataWriteDelete)
	fmt.Println(deleteValue)
}
