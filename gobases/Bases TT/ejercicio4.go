
package main

import "fmt"

func main (){

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es: %v \n", employees["Benjamin"])
	mayorDeEdad:=0

	for _,age := range employees{
		if age>21{
			mayorDeEdad+=1;
		}
	}
	fmt.Printf("Hay %v mayores de 21 \n",mayorDeEdad)

	employees["Federico"]=25
	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}