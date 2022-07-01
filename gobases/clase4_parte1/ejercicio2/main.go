package main

import (
	"errors"
	"fmt"
)

func calculoSueldo(s int) {
	if s < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	} else if s >= 150000 {
		fmt.Println(errors.New("Debe pagar impuestos"))
		return
	}
	fmt.Println("El programa finalizo OK")
}

func main() {
	var salary int = 232321321
	calculoSueldo(salary)

}
