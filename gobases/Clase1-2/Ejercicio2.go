package main

import "fmt"

func main(){
	var (
		edad int = 22
		antiguedad int = 1
		salario int =100000
	)
	
	if edad >= 22{
		if antiguedad >= 1{
			if salario >= 100000{
				fmt.Println("Se puede otorgar el credito sin intereses")
			}else{
				fmt.Println("Se puede otorgar el credito con intereses")
			}
		}else{
			fmt.Println("No puede otorgar el credito \nMotivo: No cumple con la antiguedad requerida")
		}
	}else{
		fmt.Println("No puede otorgar el credito \nMotivo: No cumple con la edad requerida")
	}
}