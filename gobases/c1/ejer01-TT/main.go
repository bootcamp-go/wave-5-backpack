package main

import "fmt"

func main() {

	palabra := "palabra"

	for i, letra := range palabra {

		fmt.Println(i, string(letra))
	}

}
