package main

import "fmt"

func main() {
	palabra := "ribonucleico"
	fmt.Printf("Longitud: %d\n", len(palabra))
	for _, elem := range palabra {

		fmt.Printf("%s \n", string(elem))
	}
}
