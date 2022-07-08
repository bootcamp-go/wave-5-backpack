package main

import (
	"fmt"
	"os"
)

type Cliente struct {
	legajo    int
	nombre    string
	dni       int
	telefono  int
	domicilio string
}

func main() {

	// Ejercicio 1

	// Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo.txt.
	// Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
	// Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).
	// Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje
	// “el archivo indicado no fue encontrado o está dañado”.
	// Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

	fmt.Println("******** EJERCICIO 1 ********")
	LeerArchivo("./customers.txt")
	fmt.Println("Ejecucion finalizada")

	//Ejercicio 2

	// 	El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
	// Legajo
	// Nombre y Apellido
	// DNI
	// Número de teléfono
	// Domicilio

	// Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos.
	// Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”.
	// Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.

	// Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un archivo .txt.
	// En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error).
	// Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:

	// 1.- generar un panic;
	// 2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.

	// Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero.
	// Esta función debe retornar, al menos, dos valores.
	// Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).

	// Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
	// “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

	fmt.Println("******** EJERCICIO 2 ********")

	fmt.Println("TAREA 1")
	informarPanicLegajo(0)

	fmt.Println("TAREA 2")
	cliente := Cliente{legajo: 1, nombre: "Denis", dni: 38797877, telefono: 15352879, domicilio: "Belgrano"}
	comprobarRepetidos(cliente, "./customers.txt")

	fmt.Println("TAREA 4")
	fmt.Println("Fin de la ejecución")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println("No han quedado archivos abiertos")

}
func LeerArchivo(path string) {

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	file, erro := os.ReadFile(path)
	if erro != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	} else {
		fmt.Println(file)
	}

}

func generarNumeroDeLegajo(numero int) int {

	if numero < 0 {
		return 0
	} else {
		return numero
	}

}

func informarPanicLegajo(legajo int) {

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	legajo = generarNumeroDeLegajo(0)
	if legajo == 0 {
		panic("Hubo un problema en la generacion del legajo")
	}

}

func comprobarRepetidos(cliente Cliente, path string) {
	LeerArchivo(path)
}
