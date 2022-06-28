package main
import "fmt"


func main(){
	
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//Busqueda de un empleado
	fmt.Println("El empleado consultado tiene", employees["Benjamin"], "años")

	//Empleados con más de 21 años
	twentyOneMore := 0
	for _, value := range employees{
		if value > 21 {
			twentyOneMore++
		}
	}
	fmt.Println("Hay ",twentyOneMore, "empleados que tienen mas de 21 años")

	//Añadiendo a Federico
	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminando a Pedro

	delete(employees, "Pedro")
	fmt.Println(employees)
	


}