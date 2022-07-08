package main

import (
	"fmt"
	"os"
)

/*Ejercicio 1 - Datos de clientes

Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt.
	1. Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica
	el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
	2. Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt”
	(recuerda lo visto sobre el pkg “os”).Dado que no contamos con el archivo necesario,
	se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar
	leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
	Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

func main() {
	defer func() {
		fmt.Println("Ejecución finalizada.")
		err := recover() // Recupera el panic - para evitar una ejecución no deseada

		if err != nil {
			fmt.Println(err)
		}
	}()

	read, err := os.ReadFile("./customer.txt")
	if err != nil {
		fmt.Println("Llegando al panic . . . 👀")
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	file := string(read)
	fmt.Println(file)
}
