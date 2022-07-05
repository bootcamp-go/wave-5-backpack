package main


import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"fmt"

)


func main() {
	archivo := file.File{Path:"tickets.csv"}
	tickets, err := archivo.Read()

	if err != nil {
		fmt.Println(err)
		return
	}
	// Funcion para obtener tickets del archivo csv
	myBookings := service.NewBookings(tickets)

	myBookings.Create(
		service.Ticket{
			Names: "Michael Jackson",
			Email:"mail@mail.com",
			Destination: "Madrid",
			Date: "15:00",
			Price: 1000,
		})

	id300, _ := myBookings.Read(300)
	fmt.Print("Este es el ticket de id 300", id300, "\n")
	id1000, _ := myBookings.Read(1000)
	fmt.Print("Este es el ticket de id 300", id1000, "\n")

	archivo.Write(myBookings.Tickets)

}
