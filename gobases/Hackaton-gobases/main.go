package main

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

const (
	READ   = "READ"
	CREATE = "CREATE"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

func main() {
	var tickets []service.Ticket
	f := &file.File{Path: "tickets.csv"}
	read, err := f.Read()
	if err != nil {
		panic(err)
	}

	//Servicio Bookings
	tickets = getTickets(read)
	s := service.NewBookings(tickets)

	var (
		OPERACION string = "CREATE"
		UpdateId  int    = 1000
		ReadId    int    = 1000
		DeleteId  int    = 1000
	)

	switch OPERACION {
	case READ:
		ticket, err := s.Read(ReadId)
		if err != nil {
			fmt.Println(err)
			return
		}
		ShowInfo(ticket)
	case CREATE:
		ticketsC := service.Ticket{
			Names:       "Harold Preciado",
			Email:       "calimalo@gmail.com",
			Destination: "Cali",
			Date:        "12:12",
			Price:       123,
		}
		ticketsC, err = s.Create(ticketsC)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(ticketsC, CREATE)
		if err != nil {
			fmt.Println(err)
			return
		}
	case UPDATE:
		tUpdate := service.Ticket{
			Names:       "Emiliano Halapeno",
			Email:       "emi@email.com",
			Destination: "Duabai",
			Date:        "20:00",
			Price:       1500,
		}
		ticketUp, err := s.Update(UpdateId, tUpdate)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(ticketUp, UPDATE)
		if err != nil {
			fmt.Println(err)
			return
		}
	case DELETE:
		id, err := s.Delete(DeleteId)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(service.Ticket{Id: id}, DELETE)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ShowInfo(ticket service.Ticket) {
	fmt.Printf("\n*** N° %d *****\n Names: %s\nEmail: %s\nDestination: %s\nDate: %s\nPrice: %d\n\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
}

func getTickets(lines [][]string) []service.Ticket {
	var tickets []service.Ticket
	for _, value := range lines {
		id, err := strconv.Atoi(value[0])
		if err != nil {
			panic(err)
		}

		price, err := strconv.Atoi(value[5])
		if err != nil {
			panic(err)
		}

		ticket := service.Ticket{
			Id:          id,
			Names:       value[1],
			Email:       value[2],
			Destination: value[3],
			Date:        value[4],
			Price:       price,
		}

		tickets = append(tickets, ticket)
	}
	return tickets
}
