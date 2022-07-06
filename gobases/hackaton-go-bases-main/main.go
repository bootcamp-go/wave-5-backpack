package main


import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"fmt"

)

const (
	READ   = "READ"
	CREATE = "CREATE"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// creo las funciones para los llamados
func create(s service.Bookings, t service.Ticket, f file.File){
	s.Create(t)
	f.Write(t)
}


func main() {
	f := file.File{Path:"./tickets.csv"}
	read, err := f.Read()

	if err != nil {
		//fmt.Println(err)
		//return

		// conviene directamente meter un panic
		panic(err)
	}

	// Funcion para obtener tickets del archivo csv
	s := service.NewBookings(read) // REVISAR XQ NO TOMA DIRECTAMENTE "read" SINO QUE HACE UN TRABAJO PREVIO

	/* s.Create(
		service.Ticket{
			Names: "Michael Jackson",
			Email:"mail@mail.com",
			Destination: "Madrid",
			Date: "15:00",
			Price: 1000,
		}) */

	var nuevoTicket = service.Ticket{
		Names: "Michael Jordan",
		Email:"mail@mail.com",
		Destination: "Chicago",
		Date: "19:00",
		Price: 3000,		
	}

	create(s, nuevoTicket, f)

	id1300, _ := s.Read(1300)
	fmt.Print("Este es el ticket de id 300", id1300, "\n")
	id1000, _ := s.Read(1000)
	fmt.Print("Este es el ticket de id 300", id1000, "\n")

	//f.Write(s.Tickets)

}
