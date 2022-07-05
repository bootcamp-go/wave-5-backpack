package main

import "fmt" 
import "strconv"

func main() {
	var temperatura int
	var humedad int
	var presion int
	temperatura=11
	humedad=80
	presion=1017
	fmt.Println("Huemdad: "+strconv.Itoa(humedad)+"%")
	fmt.Println("Temperatura: "+strconv.Itoa(temperatura)+"°C")
	fmt.Println("Presion: "+strconv.Itoa(presion)+"mb")
}
