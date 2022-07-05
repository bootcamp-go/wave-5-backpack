package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	var foundData []service.Ticket
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	dataStr := string(data)
	dataSplit := strings.Split(dataStr, "\n")
	for _, dataRow := range dataSplit {
		dataRowSplit := strings.Split(dataRow, ",")
		idTicket, err := strconv.Atoi(dataRowSplit[0])
		if err != nil {
			return nil, err
		}
		nameTicket := dataRowSplit[1]
		emailTicket := dataRowSplit[2]
		destinationTicket := dataRowSplit[3]
		dateTicket := dataRowSplit[4]
		priceTicket, err := strconv.Atoi(dataRowSplit[5])
		if err != nil {
			return nil, err
		}
		foundData = append(foundData, service.Ticket{
			Id:          idTicket,
			Names:       nameTicket,
			Email:       emailTicket,
			Destination: destinationTicket,
			Date:        dateTicket,
			Price:       priceTicket,
		})
	}
	return foundData, nil
}

func (f *File) Write(newTickets []service.Ticket) error {
	res := ""
	for i := 0; i < len(newTickets); i++ {
		ticket := newTickets[i]
		res += fmt.Sprintf("%d,%s,%s,%s,%s,%d",
			ticket.Id,
			ticket.Names,
			ticket.Email,
			ticket.Destination,
			ticket.Date,
			ticket.Price,
		)
		if i < len(newTickets)-1 {
			res += "\n"
		}
	}
	os.WriteFile(f.Path, []byte(res), 0644)
	return nil
}
