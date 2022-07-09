package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Client struct {
	Legajo         int    `json:"LEGAJO"`
	NombreCompleto string `json:"NOMBRE_COMPLETO"`
	DNI            string `json:"DNI"`
	Telefono       int    `json:"TELEFONO"`
	Domicilio      string `json:"DOMICILIO"`
}

func main() {
	/*
		solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
		Legajo
		Nombre y Apellido
		DNI
		Número de teléfono
		Domicilio

		Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.
		Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error). Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:
		1.- generar un panic;
		2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.
		Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).
		Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

		Requerimientos generales:
		Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
		Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que intenten leer archivos).
		Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la validación pertinente para el caso de error retornado).

	*/

	//Tarea 1
	var flagLegajo bool = true
	legajo, err := legajoMaker(flagLegajo)
	if err != nil {
		panic("error al generar legajo")
	}

	//Creamos un nuevo cliente
	var Customers []Client
	var cliente = Client{Legajo: legajo, NombreCompleto: "Rodrigo Gibran", DNI: "RG96", Telefono: 464, Domicilio: "Emilio Carranza"}
	//Tarea 2
	clientMaker(cliente, Customers)
	//El programa corre normalmente
	fmt.Println("main program is running...")

}
func legajoMaker(flag bool) (int, error) {
	//fmt.Println(flag)
	if flag != true {
		return 500, fmt.Errorf("un error")
	}
	ID := rand.Intn(9999)
	return ID, nil
}
func clientMaker(cliente Client, customers []Client) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("...executing recover")
		}

	}()
	//Open File
	file, err := os.OpenFile("customers.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic("err: el archivo indicado no fue encontrado o esta danhado")
	}
	//Read File
	data := make([]byte, 0)
	b, err := file.Read(data)
	if err != nil {
		fmt.Println("Read: ", err)
	}

	if b < 1 {
		fmt.Println("Empty file: ", b)
		customers = append(customers, cliente)
		clienteJSON, err := json.Marshal(customers)
		if err != nil {
			fmt.Println(err)
		}
		n, err := file.Write(clienteJSON)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("File write success %v\n", n)
		b, err = file.Read(data)
	}

	//Look if customer exist
	err1 := json.Unmarshal(data, &customers)
	if err != nil {
		fmt.Println("Unmarshall error: ", err1)
	}
	fmt.Printf("Customer LEGAJO: %v\n", customers)

	for _, v := range customers {
		if cliente.Legajo == v.Legajo {
			panic("El Cliente ya esta registrado")
		}
	}

}
