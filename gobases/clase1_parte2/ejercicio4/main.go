package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main(){
	fmt.Println("La edad de Benjamin es:",employees["Benjamin"])

	count :=0
	for _, value := range employees {
		if value > 21 {
			count++
		}
	}
	fmt.Printf("%d empleados son mayores de 21 \n",count)

	employees["Daniela"] = 20

	fmt.Println("Actualizado",employees)

	delete(employees, "Pedro")

	fmt.Println("Actualizado de nuevo",employees)
}