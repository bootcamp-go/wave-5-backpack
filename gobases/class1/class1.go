package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printMenu() {
	fmt.Println("1: Obtener prestamo")
	fmt.Println("2: ??")
	fmt.Println("0: Salir")
}

func readConsole() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func getUserData() {
	fmt.Printf("Age: ")
	age, _ := strconv.ParseInt(readConsole(), 0, 8)

	if age < 22 && age > 90 {
		fmt.Println("Not enabled")
	}

	fmt.Printf("Seniority: ")
	seniority, _ := strconv.ParseInt(readConsole(), 0, 8)

	if seniority < 1 {
		fmt.Println("Not enabled")
	}

}

func exercise2() {
	// Exercise  2
	// -----------------------------------

	var menu bool = true

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bank menu")

	for menu {
		fmt.Println("---------------------")
		printMenu()

		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("0", text) == 0 {
			fmt.Println("Good bye!")
			menu = false
		} else if strings.Compare("1", text) == 0 {
			fmt.Println("Data request")
		} else if strings.Compare("2", text) == 0 {
			fmt.Println("May the force be with you <3!")
		} else {
			fmt.Println("Please... enter the correct number -_-!")
		}

	}
}

func exercise1() {
	// Exercise  1
	// -----------------------------------
	var word string = "hola"

	fmt.Printf("Tamaño: %d\n", len(word))

	fmt.Printf("Letras: ")
	for _, letter := range word {
		fmt.Printf("%s ", string(letter))
	}
	fmt.Printf("\n")
}

func main() {

	exercise1()

	exercise2()

}
