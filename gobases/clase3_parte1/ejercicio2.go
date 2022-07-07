/*
	Ejercicio 2 - Leer archivo
	La misma empresa necesita leer el archivo almacenado, para ello requiere que: 
	se imprima por pantalla mostrando los valores tabulados, con un título 
	(tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), 
	el precio, la cantidad y abajo del precio se debe visualizar el total 
	(Sumando precio por cantidad)

	Ejemplo:

	ID                            Precio  Cantidad
	111223                      30012.00         1
	444321                    1000000.00         4
	434321                         50.50         1
							  4030062.50
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	productos, err := os.ReadFile("./productos.csv")
	
	if err != nil {
		fmt.Printf("El archivo no pudo ser leído. Error: %v", err)
	}

	for _, value := range string(productos) {
		if value == ',' {
			//fmt.Printf("%20v")
		} else if value ==';' {
			fmt.Println()
		} else {
			fmt.Printf("%20v",string(value))
		}
	}
	
}