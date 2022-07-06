// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
// Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// Luego imprimí cada una de las letras.

package main
import "fmt"

func main()  {
	var palabra string = "papanato"

	fmt.Printf("tiene esta cantidad de letras:%d\n", len(palabra))
	for i, letra := range palabra{
		fmt.Printf("posicion: %x letra: %c \n", i, letra)
	}
}