
package main

import "fmt"

func main(){

	 age:= 23
	 years_job:=0.5
	 salary:=120000

	if age<22 {
		fmt.Println("No cumple con la edad mínima")
	} else if years_job<1 {
		fmt.Println("No cumple con la cantidad de años mínimo de empleado")
		
	} else if salary<100000{
		fmt.Println("Se entrega el préstamos pero con tasa de interés pero con interés")
	} else {
		fmt.Println("Se entrega el préstamo sin interés")
	}

}
