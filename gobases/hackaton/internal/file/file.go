package file

import (
	"bufio"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open(f.Path)

	if err != nil {
		log.Fatalf("error al abrir el archivo: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	res := []service.Ticket{}
	for _, foreach := range txtlines {
		a := strings.Split(foreach, ",")
		id, _ := strconv.Atoi(a[0])
		price, _ := strconv.Atoi(a[5])
		res = append(res, service.Ticket{Id: id, Names: a[1], Email: a[2], Destination: a[3], Date: a[4], Price: price})
	}
	return res, err
}

func (f *File) Write(service.Ticket) error {
	return nil
}
