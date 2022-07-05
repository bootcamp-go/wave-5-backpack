package main

import "fmt"

func main(){
	palabra := "escultura"
	largo := len(palabra)
	fmt.Printf("largo = %v \n",largo)
	for i:=0; i<largo; i++ {
		fmt.Printf("%c \n",palabra[i])
	}
}

//a Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
//Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
//Luego imprimí cada una de las letras.
