package main

import "fmt"


func isPair(num int) {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	if (num % 2) != 0 {
		panic("el numero no es par")
	}

	fmt.Println(num, " es un número par!")
}

func main() {

	num := 3

	isPair(num)

	fmt.Println("Ejecución completada!")

}
