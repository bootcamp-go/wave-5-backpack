package main

import (
	"fmt"
	"time"
)

type prueba struct {
	edad int
}

func main() {
	p1 := prueba{5}
	p2 := prueba{4}
	list_prueba := []prueba{p1, p2}
	var puntero *prueba
	// for i, j := range list_prueba {
	// 	if i == 1 {
	// 		puntero = &j
	// 	}
	// }
	for i := 0; i < len(list_prueba); i++ {
		if i == 1 {
			puntero = &list_prueba[i]
		}
	}

	(*puntero).edad = 90
	fmt.Println(list_prueba[1])
	fmt.Println("inicio")
	currentTime := time.Now()
	fmt.Println("Current Time in String: ", currentTime.String())
	hours, minutes, _ := time.Now().Clock()
	fmt.Println("MM-DD-YYYY : ", fmt.Sprintf("%d:%02d", hours, minutes))
	fmt.Println(26 % 3)
}
