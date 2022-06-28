package main
import "fmt"

var(
	temperatura int
	humedad float32
	presion float32
)
func main(){
	temperatura = 15
	humedad = 20
	presion = 12000
	fmt.Println(temperatura, humedad, presion)
}
