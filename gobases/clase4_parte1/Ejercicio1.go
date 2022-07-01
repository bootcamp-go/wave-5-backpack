package main

import (
	"fmt"
)

type error struct {
	mensaje string
}

func (e *error) Error() string {
	return fmt.Sprintf(e.mensaje)
}

func main() {
	salary := 15000

}
