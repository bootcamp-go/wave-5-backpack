package main
import "fmt"


func main(){
	var (
		edad = 23
		esEmpleado = true
		Antiguedad = 2
		sueldo = 200000
	)

	if (edad > 22 && esEmpleado == true && Antiguedad > 1) {
		fmt.Println("Felcidades!!! Puede acceder al prestamo ")
		if sueldo > 100000 {
			fmt.Println("Tenemos buenas noticias, no se te cobraran intereses")
		}
	} else {
		fmt.Println("Lo sentinmos :( No cumple los requisitos para acceder al prestamo")
	}

}