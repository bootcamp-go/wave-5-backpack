package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	var age int = 21
	// fmt.Println("ingrese el valor de la edad inferir para calcular el numero de empleados")
	// fmt.Scanln(&age)

	var count int
	for _,v := range employees{
		if v > age{
			count++
		}
	}

	fmt.Printf("empleados mayores a 21 años son: %v \n", count)


	employees["Federico"]=25
	fmt.Println(employees)

	delete(employees,"Pedro")
	fmt.Println(employees)


}

// func countAgeUpOf(m map, age int)int{
// 	var count int
// 	for _,v := range m{
// 		if v > age{
// 			count++
// 		}
// 	}
// 	return count
// }