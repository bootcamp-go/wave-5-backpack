package main

import "fmt"

func main() {
	var (
		age          = 30
		active       = false
		yearsWorking = 2
		salary       = 5000
	)

	if age > 22 && active && yearsWorking > 1 {
		fmt.Println("Se le dará el préstamo")
		if salary > 100000 {
			fmt.Println("Y no deberá pagar intereses")
		} else {
			fmt.Println("Y deberá pagar intereses")
		}
	} else {
		fmt.Println("No se le dará el préstamo")
	}
}
