
package main

import "fmt"
import "os"

type miError struct {
	estado int
	mensaje string
}

func (err *miError) Error() string{
	return fmt.Sprintf("%d - %v",err.estado,err.mensaje)
}

func controlImpuesto(salario int)(int,error){
	if(salario<150000){
		var msj="error: el salario ingresado no alcanza el mínimo imponible"
		return 1, &miError{
			estado: salario,
			mensaje: msj,
		}
	}
	return 0,nil
}

func main(){
	salary:=50000
	_,err:=controlImpuesto(salary)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}


/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: 
el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” 
sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

*/