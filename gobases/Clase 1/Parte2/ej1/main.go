package main

import "fmt"

func main() {

	palabra := "deletreo"
	fmt.Println(len(palabra))
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}
