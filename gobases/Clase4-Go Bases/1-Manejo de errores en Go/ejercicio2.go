
package main

import (
	"fmt"
	"errors"
)
	

func main(){
	salary:=50000
	if(salary<150000){
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}


/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: 
el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” 
sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.

*/