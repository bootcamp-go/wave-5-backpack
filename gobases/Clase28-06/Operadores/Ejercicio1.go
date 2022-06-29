package main

import (
	"fmt"
)

/*
func main() {

		var palabra string
		fmt.Println("Ingrese palabra:\n ")
		fmt.Scanf("%s\n", &palabra)

		fmt.Println(len(palabra))

		for _,r := range palabra {
			fmt.Println(string(r))
		}
*/
/*
var edad int = 25
var antiguedad int = 2
var empleado bool = false
var sueldo float32 = 100500


func main(){
	if (edad > 22 && empleado && antiguedad > 1){
		fmt.Println("Cumple con los requisitos basicos para aplicar al prestamos")
		if (sueldo > 100000){
			fmt.Println("Usted esta excento de interes, por tu alto sueldo.")
		} else {
			fmt.Println("Usted tendra que pagar una tasa de interes, por su bajo salario.")
		}
	} else {
		fmt.Println("Su prestamo ha sido rechazado, no cumple con los requisitos")
	}
}
*/

/*
func main() {
	calendario := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	for key, element := range calendario {
		fmt.Println(key,element)
	}
}
*/

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])
	var cant int
	for key, _ := range employees {
		if employees[key] > 21 {
			cant++
		}
	}
	println(cant)
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}
