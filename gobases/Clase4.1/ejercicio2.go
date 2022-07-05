package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 100
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	}
}
