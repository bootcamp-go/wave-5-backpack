package main

import (
	"fmt"
	"strings"
)

func main() {
	// Ejercicio 1.
	fmt.Println("----- Exercise 1 -----")
	const name string = "Santiago"
	var SpellRaeWord = strings.Split(name, "")
	for letter := range SpellRaeWord {
		fmt.Println(SpellRaeWord[letter])
	}

	// Ejercicio 2
	fmt.Println("----- Exercise 2 -----")
	Age := 22
	IsEmployed := true
	Incomes := 990000
	if Age < 22 {
		fmt.Println("You are too young we can't borrow money")
	}
	if !IsEmployed {
		fmt.Println("You cant pay us without mensual incomes, get a job first!")
	}
	if Age >= 22 && IsEmployed {
		fmt.Println("Agree, we can borrow you money")
	}
	if Age >= 22 && IsEmployed && Incomes >= 100000 {
		fmt.Println("You don't need pay interest charges for the borrow money")
	}

	// Ejercicio 3
	fmt.Println("----- Ejercicio 3 -----")
	months := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	var monthPick int
	fmt.Scanln(&monthPick)
	var finalMonth = monthPick - 1
	fmt.Println(months[finalMonth])

	// Ejercicio 4
	fmt.Println("----- Ejercicio 4 -----")
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	for name, edad := range employees {
		if edad > 21 {
			fmt.Printf("%v es mayor de 21 \n", name)
		} else {
			fmt.Printf("%v es menor de 21 \n", name)
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
