package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println("Benjamin tiene", employees["Benjamin"], "años")
	var count int = 0
	for _, edad := range employees {
		if edad > 21 {
			count += 1
		}
	}
	fmt.Printf("Hay %+v empleados con más de 21 años \n", count)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}