package main

import (
	"fmt"
	//"reflect"
)

func main() {
	var (
		temperatura float32 = 13.8
		humedad     float32 = 80.1
		presion     float32 = 1018.0
	)
	//fmt.Println(reflect.TypeOf(main))
	fmt.Println("Temperatura:"+fmt.Sprintf("%v", temperatura)+"ºC", "Humedad:"+fmt.Sprintf("%v", humedad)+"%", "Presión:"+fmt.Sprintf("%v", presion)+"mb")
}
