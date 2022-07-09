package main

import (
	"fmt"
)

func main() {
	var myTemp float32 = 22.0
	var myHum float32 = 53.0
	var myPres float32 = 1018

	fmt.Printf("La temperatura en Salamanca es %vº mientras que la humedad es del %v porciento y la presion atmosferica de %v mib", myTemp, myHum, myPres)
}
