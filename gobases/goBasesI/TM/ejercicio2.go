package main

import "fmt"

func main() {
	var (
		temperatura int = 16
		humedad     int = 72
		presionAt   int = 1024
	)
	fmt.Printf("Clima Bogota\nTemperatura: %d%cC\nHumeda relativa: %d%c \nPresion atmosferica: %dhPa\n", temperatura, 176, humedad, 37, presionAt)
}
