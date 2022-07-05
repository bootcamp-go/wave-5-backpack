package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees)
	var contador int=0
	for key,element :=range employees {
		fmt.Printf("Nombre: %v Edad: %v\n",key,element)
		if(element>21){
			contador++
		}
	}
	fmt.Printf("Mayores de 21 años %v\n",contador)
	employees["Federico"]=25
	delete(employees,"Pedro")
	fmt.Println(employees)
}


/*Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/