package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./products.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.ReplaceAll(string(file), ",", "\t"))
}
