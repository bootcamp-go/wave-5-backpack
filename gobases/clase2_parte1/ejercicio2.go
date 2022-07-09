package main

import "fmt"


func main()  {
	notas_alumno, msg := promedios(1, 4, 5, 3, 2, 4)
	if msg != "" {
		fmt.Printf("%s\n", msg)
		return
	}
	fmt.Printf("El promedio del alumno es de: %.1f\n", notas_alumno)
}


func promedios(notas ... float32) (float32, string){
	var sum float32 = 0
	for _, nota :=  range notas{
		if nota < 0 {
			return 0, fmt.Sprintf("La nota con valor %.1f es invalida", nota)
		}
		sum += nota
	}
	return sum / float32(len(notas)), ""
}
