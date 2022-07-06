package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Telefono  int
	Domicilio string
}

//Generando id
func generarID(numberLegajo int) int {
	return numberLegajo + 1
}

func printError(value string) error {
	return fmt.Errorf("error: el campo ", value, " está vacío")
}

func crearCliente(legajoC int, nombre, apellido string, DNI, telefono int, domicilio string) (*Cliente, error) {

	emptyClient := &Cliente{}
	if legajoC == 0 {
		err := printError("legajo")
		return emptyClient, err
	}

	if nombre == "" {
		err := printError("nombre")
		return emptyClient, err
	}

	if apellido == "" {
		err := printError("apellido")
		return emptyClient, err
	}

	if DNI == 0 {
		err := printError("dni")
		return emptyClient, err
	}

	if telefono == 0 {
		err := printError("telefono")
		return emptyClient, err
	}

	if domicilio == "" {
		err := printError("domicilio")
		return emptyClient, err
	}

	cliente := &Cliente{legajoC, nombre, apellido, DNI, telefono, domicilio}
	return cliente, nil
}

func leerArchivo(fileName string) {
	//Controlo con recover para que no se aborte
	//la ejecución
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	//lectura de archivo
	infoCustomers, err := os.ReadFile(fileName)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	//Para mejorar visualización
	fmt.Println("\n¡Lectura satisfactoria!")
	fmt.Println("--------------------------")
	fmt.Print("Info de txt customers\n--------------------------\n", string(infoCustomers), "\n")
}

func printInfo(cliente *Cliente) {
	data, _ := os.ReadFile("./customers.txt")

	returnValue := string(data)
	value := fmt.Sprint(cliente.Legajo, "\t", cliente.Nombre, "\t", cliente.Apellido, "\t", cliente.DNI, "\t", cliente.Telefono, "\t", cliente.Domicilio, "\n")
	returnValue += value

	//WriteFile para escribirlo en el txt creado
	dataProducto := []byte(returnValue)
	_ = os.WriteFile("./customers.txt", dataProducto, 0644)
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	legajo := r.Int()
	//Tarea 1
	numberLegajo := generarID(legajo)

	var (
		nombre    string
		apellido  string
		dni       int
		telefono  int
		domicilio string
	)

	nombre = "Martha"
	apellido = "Hernandez"
	dni = 3123425
	telefono = 1232233
	domicilio = "Calle 26"

	leerArchivo("./customers.txt")

	newclient, err := crearCliente(numberLegajo, nombre, apellido, dni, telefono, domicilio)
	if err != nil {
		fmt.Println(err)
	} else {
		printInfo(newclient)
		fmt.Println("¡Cliente creado!\n ----------\nInfomación nuevo cliente\n----------\n")
		fmt.Println("Legajo", newclient.Legajo)
		fmt.Println("Nombre", newclient.Nombre)
		fmt.Println("Apellido", newclient.Apellido)
		fmt.Println("DNI", newclient.DNI)
		fmt.Println("telefono", newclient.Telefono)
		fmt.Println("Domicilio", newclient.Domicilio)
	}

	//Probando manualmente el ingreso de un cliente al txt
	//returnValue := ""
	clienteExam := &Cliente{1, "Luz", "Carime", 22334273, 27738783, "Carrera 21"}

	printInfo(clienteExam)
}
