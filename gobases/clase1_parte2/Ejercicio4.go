package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es :",employees["Benjamin"])

	var mayores int
	for _, e := range employees{
		if e > 21 {
			mayores ++
		}
	}
	fmt.Println("Los mayores a 21 años son :",mayores)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)


}