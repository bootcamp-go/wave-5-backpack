package main

import "fmt"

func main() {

	var palabra string = "Golang"

	fmt.Println(len(palabra))
	for _, v := range palabra {
		fmt.Printf("%c\n", v)
	}

}
