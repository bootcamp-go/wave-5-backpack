package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		if len(values) != 6 {
			err = fmt.Errorf("error al separar los campos en la linea %d. Deberian de resultar 6 campos y eran %d", i, len(values))
			continue
		}
		id, err1 := strconv.ParseInt(values[0], 10, 32)
		if err1 != nil {
			err = fmt.Errorf("error al obtener el id, no se pudo parsear a int. Linea: %d", i)
			continue
		}
		names, email, destination, date := values[1], values[2], values[3], values[4]
		precio, err2 := strconv.ParseInt(values[5], 10, 32)
		if err2 != nil {
			err = fmt.Errorf("error al obtener el precio, no se pudo parsear a int. Linea: %d", i)
			continue
		}
		tickets = append(tickets, service.Ticket{Id: int(id), Names: names, Email: email, Destination: destination, Date: date, Price: int(precio)})
		i++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tickets, err
}

func (f *File) Write(tickets []service.Ticket) error {
	resultado := ""
	for _, value := range tickets {
		resultado += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", value.Id, value.Names, value.Email, value.Destination, value.Date, value.Price)
	}
	err := os.WriteFile(f.Path, []byte(resultado), 0644)
	return err
}
