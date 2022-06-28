package main
import "fmt"
func main(){
	var temperatura float32
	var humedad int
	var presion float32

	temperatura, humedad, presion = 32.1, 50, 2.65
	fmt.Println(temperatura, humedad, presion)
}