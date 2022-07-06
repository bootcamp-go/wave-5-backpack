package main

import(
	"fmt"
	"errors"
)

func main(){
	promedio, err := calcularPromedio(1.0,4.5,5.0,3.3,2.0)
	if err != nil{
		fmt.Println("¡Error!.",err)
	}else{
		fmt.Printf("Promedio del alumno: %.2f\n",promedio)
	}
}

func calcularPromedio(calificaciones ...float64) (float64, error){
	//promedio para suma de cantidades
	//cantidaddato es la cantidad de calificaciones ingresadas
	//parsecantidad para parseo de int a float para obtener promedio exacto

	var(
		promedio = 0.0
		cantidadDato = len(calificaciones)
		parseCantidad = float64(cantidadDato)
	)
	for _, calificacion := range calificaciones {
		//error en caso e negativo
		if calificacion < 0 {
			return 0, errors.New("Una de las calificaciones es negativa")
		}
		promedio += calificacion
	}

	//cálculo de promedio directamente en return
	return (promedio/parseCantidad), nil
}

