package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	defer fmt.Println("ejecución finalizada")

	data, err := os.ReadFile("./gobases/4/tt/customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	r := strings.NewReader(string(data))
	b, err := io.ReadAll(r)

	fmt.Println(string(b))
}
