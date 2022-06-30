package main

import "fmt"

func main() {
	nombre, edad := "Kim", 22

	res := fmt.Sprint(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Println(res)

	res = fmt.Sprintf("%s tiene %d años de edad.\n", nombre, edad)
	fmt.Println(res)
}
