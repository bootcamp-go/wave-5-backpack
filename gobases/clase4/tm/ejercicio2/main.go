package main

import (
	"errors"
	"fmt"
)

func main() {
	paga := 160000

	err := checkImpuesto(paga)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	noPaga := 100000

	err = checkImpuesto(noPaga)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func checkImpuesto(salary int) error {
	if salary < 150000 {
		return errors.New("No alcanza el mínimo no imponible")
	}

	return nil
}
