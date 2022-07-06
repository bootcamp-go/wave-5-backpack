package main

import "fmt"

func main(){
	var(
		edad int = 26
		antiguedad int = 1
		sueldo float32 = 1000000
	)

	if edad > 22 && (antiguedad > 1 && antiguedad != 0) {
		if(sueldo > 100000){
			fmt.Println("Aprovado. No se le cobra interés")
		}else{
			fmt.Println("Aprovado con interés")
		}
	}else{
		fmt.Println("No aprovado")
	}
}