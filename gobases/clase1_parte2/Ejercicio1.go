package main
import "fmt"

func main(){
	var palabra string = "palabra"
	fmt.Printf("Cantidad de letras = %v \n", len(palabra))
	for _, letra := range palabra {
		fmt.Printf("%s \n", string(letra))
	}
	// for i := 0; i < 7; i++{
	// 	fmt.Printf("%d \n", palabra[i])
	// }
}

//La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
// Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// Luego imprimí cada una de las letras.
