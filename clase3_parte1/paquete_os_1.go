package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Setenv("NAME", "gopher")

	fmt.Printf("Error: %v\n", err)

	value := os.Getenv("NAME")

	fmt.Printf("Variable de entorno: %s\n", value)

	value, ok := os.LookupEnv("USERNAME")

	fmt.Printf("Varaible de enotrno: %s| Existe?: %t\n", value, ok)

	os.Exit(1)
}
