package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	var fileRead []service.Ticket
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	} else {
		dataString := string(data)
		dataList := strings.Split(dataString, "\n")
		for _, l := range dataList {
			if l != "" {
				lSplit := strings.Split(l, ",")
				if len(lSplit) != 6 {
					continue
				}
				id, msg := strconv.Atoi(lSplit[0])
				if msg != nil {
					continue
				}
				names := lSplit[1]
				if names == "" {
					continue
				}
				email := lSplit[2]
				if email == "" {
					continue
				}
				destino := lSplit[3]
				if destino == "" {
					continue
				}
				date := lSplit[4]
				if date == "" {
					continue
				}
				price, msg := strconv.Atoi(lSplit[5])
				if msg != nil {
					continue
				}
				ticket := service.Ticket{
					Id:          id,
					Names:       names,
					Email:       email,
					Destination: destino,
					Date:        date,
					Price:       price,
				}
				//fmt.Println(ticket)
				fileRead = append(fileRead, ticket)
			}
		}
	}
	return fileRead, nil
}

func (f *File) Write(t service.Ticket) error {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil
	}
	dataString := string(data)
	ticket := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d\n", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	dataString += ticket
	dataBit := []byte(dataString)
	err = os.WriteFile(f.path, dataBit, 0644)

	return err
}

func (f *File) WriteAll(tickest []service.Ticket) error {
	var dataString string
	for _, t := range tickest {
		dataString += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	}
	dataBit := []byte(dataString)
	err := os.WriteFile(f.path, dataBit, 0644)
	return err
}

func (f *File) ChangePath(str string) {
	f.path = str
}

//func main() {
//	ticket := service.Ticket{Id: 1005, Names: "Cristian Velez", Email: "cristian.velez@", Destination: "Bogota", Date: "13:30", Price: 390}
//	archivo := File{"../../tickets.csv"}
//	err2 := archivo.Write(ticket)
//	fmt.Println(err2)
//	tickets, err := archivo.Read()
//	fmt.Println(err)
//	fmt.Println(len(tickets))
//}
