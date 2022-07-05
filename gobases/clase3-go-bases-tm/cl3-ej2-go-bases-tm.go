/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Leer Archivo
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		The same company needs to read the stored file, for this it requires that:
		it is printed on the screen showing the tabulated values, with a title screen
		showing the tabulated values, with a title (tabbed on the left for the ID and
		on the right for the Price and Quantity), the price and on the right for
		the Price and Quantity), the price, the quantity and below the price the
		total should be displayed (adding price by quantity).

		Example:
			ID                            Precio  Cantidad
			111223                      30012.00         1
			444321                    1000000.00         4
			434321                         50.50         1
			4030062.50

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//	CONSTANTS
const (
	coma  = ","
	space = " "
)

//	STRUCT : Producto
type Producto struct {
	ID       int
	PRECIO   float64
	CANTIDAD int
}

//	FUNCTION : imprimirArchivo
func imprimirArchivo(data []byte) {
	var total float64

	filas := strings.Split(string(data), "\n")
	for _, col := range filas {

		palabra := strings.Split(col, ",")

		for i, c := range palabra { // index, character
			/*| Columna 0 = Case 0 | Columna 1 = Case1 | Columna 2 = Case2 |*/
			switch i {
			case 0: // Primera Columna
				fmt.Printf("%12s\t", c)
			case 1: // Segunda Columna
				digit, err := strconv.ParseFloat(c, 64)
				if err != nil {
					fmt.Printf("%12s\t", c)
				} else {
					total += digit
					fmt.Printf("%12.2f\t", digit)
				}
			case 2: // Tercera Columna
				fmt.Printf("%12s\t", c)
			}
		}
		fmt.Println("")
	}
	// Total de Precios
	fmt.Printf("\tTOTAL \t%12.2f\n", total)
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t\t|| Leer Archivo ||")

	data, err := os.ReadFile("./productos.csv")
	if err != nil {
		fmt.Printf("Error lectura directorio: %v", err)
	}
	imprimirArchivo(data)
}
