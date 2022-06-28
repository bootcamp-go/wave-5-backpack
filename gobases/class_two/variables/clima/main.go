package main 
import "fmt"


var(
	temperatura int
	humedad float32
	presion float32
)

func main(){
	temperatura= 15
	humedad = 20
	presion= 120025

	fmt.Println( "Hoy estamos a:", temperatura, "ºC con una humedad de: ", humedad, " y una presion:", presion)

	fmt.Printf(" Hoy estamos a %d C con una humedad de %v y una presión de: %v", temperatura, humedad, presion)
}