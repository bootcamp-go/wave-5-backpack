/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. 
1) Se solicita generar una función en la cual 
a) se le pueda pasar N cantidad de enteros 
b) y devuelva el promedio 
c) y un error en caso que uno de los números ingresados sea negativo

*/

/*
package main


import (
	"fmt"
	"errors"
)



func main() {

	//res, err := calcularPromedio(1,5,10,7, -1) 
	// se espera 4,4 y error

	res, err := calcularPromedio(1,5,10,7,  2) 
	// se espera 25
	

	fmt.Println("El promedio es: ", res)	
	if err != nil {
	// hubo error
	fmt.Println(err)	
	}
}

func calcularPromedio (calificaciones ... int) (float64,  error) {

	var totalSlice float64 = 0.0
	totalSlice = float64(len(calificaciones))
	sumaTotal := 0
	hayErrores := false
	for _, valor := range calificaciones {
	if valor < 0 {
		hayErrores = true
	}
	sumaTotal += valor
	}

	var promedio float64 = 0.0 
	promedio = float64(sumaTotal) / totalSlice
	if hayErrores {
	err := errors.New("Se encontro un numero negativo")
	return promedio,err
	}else{
		return promedio, nil
	}
	
}


*/