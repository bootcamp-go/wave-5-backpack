package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso invalido, debe incluir accion y datos")
		os.Exit(1)
	}

	var ticket service.Ticket

	data := os.Args[2]
	err := json.Unmarshal([]byte(data), &ticket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tickets []service.Ticket
	path := "tickets.csv"

	f := file.NewFile(path)
	tickets, err = f.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := service.NewBookings(tickets)

	switch os.Args[1] {

	case "create":
		new_ticket, err := b.Create(ticket)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tickets = append(tickets, new_ticket)
		f.Write(tickets)

		ticket_json, err := json.MarshalIndent(new_ticket, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Ticket creado:")
		fmt.Println(string(ticket_json))

	case "read":
		read_ticket, err := b.Read(ticket.Id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ticket_json, err := json.MarshalIndent(read_ticket, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Ticket solicitado:")
		fmt.Println(string(ticket_json))

	case "update":
		updated_ticket, err := b.Update(ticket.Id, ticket)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		first := tickets[:ticket.Id-1]
		first = append(first, updated_ticket)
		last := tickets[ticket.Id:]

		tickets = append(first, last...)
		f.Write(tickets)

		ticket_json, err := json.MarshalIndent(updated_ticket, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Ticket actualizado:")
		fmt.Println(string(ticket_json))

	case "delete":
		deleted_ticket, err := b.Delete(ticket.Id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		first := tickets[:deleted_ticket-1]
		last := tickets[deleted_ticket:]

		tickets = append(first, last...)
		f.Write(tickets)

		fmt.Printf("Ticket de id %d borrado\n", deleted_ticket)

	}

	fmt.Println("Fin de ejecución")
}
