package main
import "fmt"


func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int = 0
	fmt.Printf("La edad de Benjamín es: %d \n", employees["Benjamin"])
	for _, edad := range employees{
		if edad >= 21 {
			mayores += 1
		}
	}
	fmt.Printf("Empleados mayores de 21 años: %d \n", mayores)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Printf("%v \n", employees)
}

// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 

// Por otro lado también es necesario: 
// -Saber cuántos de sus empleados son mayores de 21 años.
// -Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// -Eliminar a Pedro del mapa.
