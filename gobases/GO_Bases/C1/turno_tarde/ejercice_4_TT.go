package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de benjamis es %d\n", employees["Benjamin"])
	var age_employees21 []string
	for name, employee := range employees {
		if employee > 21 {
			age_employees21 = append(age_employees21, name)
		}
	}
	fmt.Println("Los empleados mayores a 21 son: \n", age_employees21)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
