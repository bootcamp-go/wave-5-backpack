package file

import "https://github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	return nil, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
