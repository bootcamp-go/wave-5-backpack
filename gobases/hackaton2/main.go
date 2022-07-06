package main

import (
	"fmt"
	"hackaton-go-bases/internal/service"
	"os"

	"hackaton-go-bases/internal/file"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv

	path := "./tickets.csv"
	file := file.File{Path: path}
	tickets, err := file.Read()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	funcs := service.NewBookings(tickets)

	fmt.Println("")
	fmt.Println("PRUEBAS CON DATOS INVALIDOS")
	t1 := service.Ticket{Names: "", Email: "jessica@gmail.com", Destination: "", Date: "18:30", Price: 520}
	fmt.Println(funcs.Create(t1))
	fmt.Println(funcs.Read(1010))
	t1 = service.Ticket{Names: "Paola", Email: "paola@gmail.com", Destination: "Canadá", Date: "20:30", Price: 800}
	fmt.Println(funcs.Update(1010, t1))
	fmt.Println(funcs.Delete(1010))

	fmt.Println("")
	fmt.Println("PRUEBAS CON DATOS VALIDOS")
	t1 = service.Ticket{Names: "Jessica", Email: "jessica@gmail.com", Destination: "Colombia", Date: "18:30", Price: 520}
	fmt.Println(funcs.Create(t1))
	fmt.Println(funcs.Read(1001))
	t1 = service.Ticket{Names: "Paola", Email: "paola@gmail.com", Destination: "Canadá", Date: "20:30", Price: 800}
	fmt.Println(funcs.Update(1001, t1))
	fmt.Println(funcs.Read(1001))

	// Eliminando el registro con ID 1000
	//funcs.Delete(1000)

	// Escritura del archivo sin el registro con ID 1000 y con el nuevo registro con ID 1001
	//file.Write(funcs.GetTickets())

}
