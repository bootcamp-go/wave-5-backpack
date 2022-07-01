package main

//%10.2f <-- ejemplo de tabulacion
import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("./Archivo.txt")
	if err != nil {
		fmt.Printf("Error lectura: %v", err)
	}
	//fmt.Printf("file: %v \n", string(data))
	lect := strings.Split(string(data), ",")
	fmt.Println(lect)

}
