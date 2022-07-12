package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayoresDe21 int = 0
	fmt.Println("")
	fmt.Printf("Benjamin tiene %d años\n",employees["Benjamin"])
	
	fmt.Println("")
	for _, element := range employees{
		if element > 21 {
			mayoresDe21++
		}
	}
	fmt.Printf("Hay %d empleados mayores de 21 años\n",mayoresDe21)
	
	employees["Federico"] = 21

	delete(employees, "Pedro")

}