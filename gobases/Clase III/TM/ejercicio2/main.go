package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leerArchivo(text string) {

	archivoOrdenado := strings.Split(text, "\n")
	totalProductos := 0.0

	for i := 0; i < len(archivoOrdenado); i++ {
		archivoSinComas := strings.Split(archivoOrdenado[i], ",")
		if i == 0 {
			textFinal := fmt.Sprintf("%v \t\t\t%v \t %v", archivoSinComas[0], archivoSinComas[1], archivoSinComas[2])
			fmt.Println(textFinal)
		} else {
			f1, _ := strconv.ParseFloat(archivoSinComas[1], 64)
			f2, _ := strconv.ParseFloat(archivoSinComas[2], 64)
			totalProductos += f1 * f2
			textFinal := fmt.Sprintf("%v %23.2f  %9.0f", archivoSinComas[0], f1, f2)
			fmt.Println(textFinal)
		}
	}

	resultado := fmt.Sprintf("%30.2f", totalProductos)
	fmt.Println(resultado)
}

func main() {

	data, _ := os.ReadFile("./archivo.csv")
	leerArchivo(string(data))

}
