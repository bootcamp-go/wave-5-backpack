package main

import "fmt"

func main() {
	//Funciones en GO
	//1
	fmt.Printf("Impuesto a pagar: %.2f\n", impuestosDeSalario(200000))
	//2
	promedio, err := calcularPromedio(10, 2, 3, 4, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio de notas:", promedio)
	}

	//3
	fmt.Println("Salario:", calcularSalario(60, "A"))

	//4
	calcularEstadisticas(MAXIMUM)

	//5
	calcularCantidadDeAlimento("perro")
}