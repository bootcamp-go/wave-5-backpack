package main
import "fmt"


var(
	temperatura int
	humedad float64
	presion float64
)
	
func main () {
	temperatura = 15 
	humedad = 0.60
	presion = 1027

	fmt.Println("temperatura", temperatura)
	fmt.Println("humedad", humedad)
	fmt.Println("presion", presion)
}