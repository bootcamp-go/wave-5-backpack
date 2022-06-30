package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	r := strings.NewReader("some io.Reader stream to be read\n")

	bytesEscritos, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Printf("Error en copia: %v", err)
	}
	fmt.Printf("Bytes escritos %d\n", bytesEscritos)

	arrayBytes, err := io.ReadAll(r)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", string(arrayBytes))

	io.WriteString(os.Stdout, "Hello world!\n")
}
