package main
import "fmt"


func main(){
	mes := 20

	switch {
	case mes == 1:
		fmt.Println("Enero")
	case mes == 2:
		fmt.Println("Febrero")
	case mes == 3:
		fmt.Println("Marzo")
	case mes == 4:
		fmt.Println("Abril")
	case mes == 5:
		fmt.Println("Mayo")
	case mes == 6:
		fmt.Println("Junio")
	case mes == 7:
		fmt.Println("Julio")
	case mes == 8:
		fmt.Println("Agosto")
	case mes == 9:
		fmt.Println("Septiembre")
	case mes == 10:
		fmt.Println("Octubre")
	case mes == 11:
		fmt.Println("Noviembre")
	case mes == 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("El valor no corresponde a un mes")
	}
	

}