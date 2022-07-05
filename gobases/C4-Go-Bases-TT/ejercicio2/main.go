package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Ejercicio 2 - Registrando clientes

// El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
// Los datos requeridos para registrar a un cliente son:
//   - Legajo
//   - Nombre y Apellido
//   - DNI
//   - Número de teléfono
//   - Domicilio

// * Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos.
//   Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”.
//   Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.

// * Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe.
//   Para ello, necesitas leer los datos de un archivo .txt.
//   En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt”
//   (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error).
//   Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:
//      1.- generar un panic;
//      2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”,
//          y continuar con la ejecución del programa normalmente.

// * Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar
//   que todos los datos a registrar de un cliente contienen un valor distinto de cero.
//   Esta función debe retornar, al menos, dos valores.
//   Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero
//   (recuerda los valores cero de cada tipo de dato, ej: 0, "", nil).

// * Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
//   "Fin de la ejecución", "Se detectaron varios errores en tiempo de ejecución" y "No han quedado archivos abiertos" (en ese orden).
//   Utiliza defer para cumplir con este requerimiento.

// Requerimientos generales:
//   * Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
//   * Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error
//     (por ejemplo las que intenten leer archivos).
//     Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello
//     (realiza también la validación pertinente para el caso de error retornado).

const FILE string = "customers.txt"

type customers struct {
	legajo    string
	nombre    string
	apellido  string
	dni       string
	telefono  string
	domicilio string
}

func generarLegajo() *string {
	prefix := "CL"
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(100)
	if rand%2 == 0 {
		leg := prefix + "-" + strconv.Itoa(rand)
		return &leg
	}
	return nil
}

func readFile() []customers {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Print("\n", err, "\n")
		}
	}()

	// Leemos el archivo txt
	data, err := os.ReadFile("./" + FILE)

	// Verificamos que se pueda leer el archivo
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}

	return readData(data)
}

func readData(bytes []byte) []customers {
	data := []customers{}
	legajo, nombre, apellido, dni, telefono, domicilio := "", "", "", "", "", ""

	// Separamos la información leída por saltos de linea
	lineas := strings.Split(string(bytes), "\n")

	// Verificamos la información por cada línea leída
	for _, p := range lineas {
		// Separamos cada línea por comas
		linea := strings.Split(p, ",")

		// Se verifica que cada linea tengan la misma cantidad de datos
		if len(linea) == 6 {
			for i, l := range linea {
				switch i {
				case 0:
					legajo = l
				case 1:
					nombre = l
				case 2:
					apellido = l
				case 3:
					dni = l
				case 4:
					telefono = l
				case 5:
					domicilio = l
				}
			}

			// Asignamos la información al arreglo
			custom := customers{legajo: legajo, nombre: nombre, apellido: apellido, dni: dni, telefono: telefono, domicilio: domicilio}
			data = append(data, custom)
		}
	}

	return data
}

func validateInfo(legajo string, nombre string, apellido string, dni string, telefono string, domicilio string) (customers, error) {
	custom := customers{}
	if nombre == "" || apellido == "" || dni == "" || telefono == "" || domicilio == "" {
		return custom, errors.New("Toda la información es obligatoria!")
	}
	custom = customers{legajo: legajo, nombre: nombre, apellido: apellido, dni: dni, telefono: telefono, domicilio: domicilio}
	return custom, nil
}

func searchByDNI(c *[]customers, dni string) bool {
	for _, v := range *c {
		if strings.Compare(v.dni, dni) == 0 {
			return true
		}
	}
	return false
}

// Función para guardar los clientes en el archivo txt
func saveCustomer(customers *[]customers) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	customersData := generateData(customers)
	// Guardamos el archivo en disco
	customersByte := []byte(customersData)
	err := os.WriteFile("./"+FILE, customersByte, 0644)

	// Si ocurrio un error lo mostramos al usuario
	if err != nil {
		panic("error: no se puede guardar el archivo\n")
	}

	fmt.Print("\nSe ha guardado el archivo!")
}

// Función para generar la cadena de texto del txt
func generateData(p *[]customers) string {
	// Cadena de texto para guardar la información del txt
	data := ""

	// Generamos la información en formato CVS para ser guardada en disco
	for _, c := range *p {
		data += fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", c.legajo, c.nombre, c.apellido, c.dni, c.telefono, c.domicilio)
	}

	return data
}

func main() {
	fmt.Println("Ejercicio 2 - Registrando clientes")

	var legajo *string
	nombre := ""
	apellido := ""
	dni := ""
	telefono := ""
	domicilio := ""

	legajo = generarLegajo()
	if legajo == nil {
		panic("No se puede generar el Legajo")
	} else {
		customersFile := readFile()

		fmt.Print("\nIngrese Nombre: ")
		fmt.Scanf("%s", &nombre)
		fmt.Print("Ingrese Apellido: ")
		fmt.Scanf("%s", &apellido)
		fmt.Print("Ingrese DNI: ")
		fmt.Scanf("%s", &dni)
		fmt.Print("Ingrese Teléfono: ")
		fmt.Scanf("%s", &telefono)
		fmt.Print("Ingrese Domicilio: ")
		fmt.Scanf("%s", &domicilio)

		customer, err := validateInfo(*legajo, nombre, apellido, dni, telefono, domicilio)
		if err != nil {
			fmt.Println(err)
		} else if len(customersFile) > 0 {
			if searchByDNI(&customersFile, customer.dni) {
				fmt.Print("\nYa existe el cliente")
			} else {
				// Agregamos el cliente al archivo customers.txt
				customersFile = append(customersFile, customer)
				saveCustomer(&customersFile)
			}
		} else {
			// Creamos el archivo customers.txt y agregamos el cliente
			customersFile = append(customersFile, customer)
			saveCustomer(&customersFile)
		}

		fmt.Print("\nFin de la ejecución\n\n")
	}
}
