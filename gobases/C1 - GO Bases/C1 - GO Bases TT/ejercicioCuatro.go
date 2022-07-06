package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	//Imprimir edad de benjamin
	fmt.Printf("Benjamin tiene %d años \n",employees["Benjamin"])

	//Cuántos amyores a 21 años
	suma := 0
	for _,element  := range employees {
		if element > 21 {
			suma++
		}
	}
	fmt.Println("Cantidad de empleados mayores a 21 años:",suma)

	//Agregar empleado nuevo
	employees["Federico"] = 25
	fmt.Println("Empleados con nuevo empleado:\n",employees)

	//Eliminar a Pedro del mapa
	delete(employees, "Pedro")
	fmt.Println("Map con Pedro eliminado\n",employees)
}