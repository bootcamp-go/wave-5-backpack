package main

import "fmt"

func main() {
	var palabra = "bootcampGo!"
	fmt.Println(len(palabra))

	for _, char := range palabra {
		fmt.Println(string(char))
	}
}
