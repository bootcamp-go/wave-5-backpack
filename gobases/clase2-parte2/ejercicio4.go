package main
import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, 
"Darío": 44, "Pedro": 30}
var counter = 0

func main(){
	fmt.Println(employees["Benjamin"])

	for _, edad := range employees{
		if edad > 21{
			counter++
		}
	}
	fmt.Println("La cantidad de empleados que son mayores de 21" , 
counter)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
