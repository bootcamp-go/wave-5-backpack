package file

import (
	"fmt"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"os"
	"strings"
	"strconv"
)

type File struct {
	Path string
}

func (f File) Read() ([]service.Ticket, error) {
	archivo, err := os.ReadFile(f.Path)
	if err !=nil{
		fmt.Println("hay un error")
		return nil, err
	} 
	var data []service.Ticket
	arch := strings.Split(string(archivo), "\n")
	
	for _,linea :=range arch {
		lineaSplit := strings.Split(linea,",")

		id,_ := strconv.Atoi(lineaSplit[0])
		name := lineaSplit[1]
		mail := lineaSplit[2]
		dest := lineaSplit[3]
		date := lineaSplit[4]
		price,_ := strconv.Atoi(lineaSplit[5])
		var ticket = service.Ticket{Id: id, Names: name, Email: mail, Destination: dest , Date: date, Price: price} 
		data = append(data, ticket)
	}
	return data, nil
}

/* func pasarAString(bookings []service.Ticket) string {
	
	var datos string
	for _, b := range bookings {
		datos += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n",b.Id,b.Names,b.Email,b.Destination,b.Date,b.Price)
	}

	datos = datos[:len(datos)-1]

	return datos
} */

func pasarAString(t service.Ticket) string {
	
	var datos string
	datos = fmt.Sprintf("%d,%s,%s,%s,%s,%d",t.Id,t.Names,t.Email,t.Destination,t.Date,t.Price)

	return datos
}

func (f *File) Write(t service.Ticket) error {
	ticketToAdd := []byte(pasarAString(t))

	data, noData := os.ReadFile(f.Path)
	if noData != nil {
		return noData
	}

	data = ticketToAdd

	err := os.WriteFile(f.Path, data, 0644)
	
	if err != nil {
		return err
	}
	return nil
}


