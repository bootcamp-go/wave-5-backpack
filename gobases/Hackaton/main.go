package main 

import (
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/tree/tapia_luis/gobases/Hackaton/internal/service"
	"github.com/bootcamp-go/wave-5-backpack/tree/tapia_luis/gobases/Hackaton/internal/file"
)

const(
	create = "create" 
	read = "read"
	update = "update"
	delete = "delete"
)

// Id                              int
// Names, Email, Destination, Date string
// Price                           int
func main(){
	f := file.File{Path:"./tickets.csv"}
	ticketLs,err1 := f.Read()
	if err != nil{
		panic(err)
	}
	book := service.NewBookings(ticketLs)
	newTicket := service.Ticket{Names:"Luis Tapia",Email:"luis@mail.com",Destination:"chile",Date:"20:20",Price:20}
	a,err2 := book.Create(newTicket)
	if err2 != nil{
		panic(err)
	}
	fmt.Println(a)
	// A := service.Ticket{1,"luis","mail@","arg","11:11",20}
	// fmt.Println(A)
	// B := file.File{".path"}
	// fmt.Println(B)
}