// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla
//mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
// el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

// Ejemplo:

// ID                            Precio  Cantidad
// 111223                      30012.00         1
// 444321                    1000000.00         4
// 434321                         50.50         1
//                           4030062.50
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./../ejercicio1/productos.csv")
	var cadena, cadena2 string
	var cadenaAux []string
	var totalPrecio float64

	cadena = string(data)

	cadena2 = strings.ReplaceAll(cadena, "\n", ",")
	cadenaAux = strings.Split(cadena2, ",")

	fmt.Println("ID\t\t\tPRECIO\t\t\tCANTIDAD")
	fmt.Println(strings.ReplaceAll(cadena, ",", "\t\t\t"))

	totalPrecio = calcularTotal(cadenaAux)

	fmt.Println("TOT:\t\t\t", totalPrecio)

}

func calcularTotal(c []string) float64 {
	var aux int = 1
	var precioTotal float64

	for i, val := range c {

		var total float64
		total, _ = strconv.ParseFloat(val, 64)
		if i == 1 || i == aux {
			precioTotal += total
			aux = i + 3
		}

	}
	return precioTotal

}
