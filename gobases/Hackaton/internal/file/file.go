package file

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/tapia_luis/gobases/Hackaton/internal/service"
	"fmt"
	"os"
	"strings"
	"strconv"
)
const(
	create = "create" 
	read = "read"
	update = "update"
	delete = "delete"
)
type File struct {
	Path string
}

func (f File) Read() ([]service.Ticket, error) {
	fmt.Println(f)
	data, err := os.ReadFile(f.Path)
	if err != nil{
		return nil,fmt.Errorf("Ah ocurrido un error: %v",err)
	}
	fileText := string(data) 
	rows := strings.Split(fileText,"\n")
	ts:= []service.Ticket{}
	for _,r := range rows{
		cols := strings.Split(r,",")
		t := service.Ticket{}
		for key,c := range cols {
			switch key {
				case 0:
					intId,_ := strconv.Atoi(c)
					t.Id = intId
				case 1:
					t.Names = c
				case 2:
					t.Email = c
				case 3:
					t.Destination = c
				case 4:
					t.Date = c
				case 5:
					intPrice,_ := strconv.Atoi(c)
					t.Price = intPrice
			}
		}
		ts = append(ts,t)
	}
	fmt.Println(ts)
	return ts,nil
}
