package main
import "fmt"
/* import "strconv" */


func main(){
	palabra := "cosito"

	for _, letra := range palabra{
		/* fmt.Println(strconv.QuoteRune(letra)) */
		fmt.Println(string(letra))
	}
	fmt.Println("El numero de letras es: ", len(palabra))
}