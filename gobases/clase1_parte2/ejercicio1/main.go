package main

import "fmt"

func main() {
	palabra := "academia"

	fmt.Printf("Largo de la palarbra: %v\n", len(palabra))

	for _, v := range palabra {
		fmt.Printf("%s\n", string(v))
	}
}
