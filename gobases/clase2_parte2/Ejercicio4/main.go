package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {

	fmt.Println("La edad de Benjamin es: ", (employees["Benjamin"]))
	employees["Federico"] = 25
	fmt.Println("Se agrego un a federico ", employees)
	delete(employees, "Pedro")
	fmt.Println("Se Elimino a Pedro ", employees)
}
