/*
Realizar una aplicación que contenga una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?


*/
package main


import "fmt"


func main() {

	var meses = map[int]string{1: "Enero",2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	var mes int = 12

	fmt.Println("Mes seleccionado:", meses[mes])	

    }


