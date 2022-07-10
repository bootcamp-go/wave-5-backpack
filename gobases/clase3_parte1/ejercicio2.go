package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	read, err := os.ReadFile("./productos.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(read)
	fmt.Println(strings.ReplaceAll(data, ";", "\t\t\t"))
}
