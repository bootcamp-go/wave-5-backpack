package main

import(
	"fmt"
	"os"
	"strings"
)

type Cliente struct{
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Telefono  int
	Domicilio string
}

//Generando id
func generarID(valueId *int){
	*valueId++
}

func crearCliente(legajoC int, nombre, apellido string, DNI, telefono int, domicilio string) string{
	cliente := Cliente{legajoC, nombre, apellido, DNI, telefono, domicilio}
	stringCliente := fmt.Sprintf("%d\t%s\t%s\t%d\t%d\t%s\n", cliente.Legajo, cliente.Nombre, cliente.Apellido,cliente.DNI,cliente.Telefono,cliente.Domicilio)
	return stringCliente
}

func main(){
	var legajo int = 3
	//Tarea 1
	generarID(&legajo)

	//Tarea 2
	returnValue := ""
	dataCustomer, err := os.ReadFile("./customers.txt")

	returnValue += string(dataCustomer)
	returnValue += fmt.Sprint(legajo, "Martha", "Hernandez", 3123425, 1232233, "Calle 26")
	
	//WriteFile para escribirlo en el txt de clientes
	dataCliente := []byte(returnValue)
	errData := os.WriteFile("./customers.txt", dataCliente, 0644)
	if errData != nil {
		fmt.Println(errData)
	}

	splitJump := strings.Split(string(dataCustomer), "\n")

	for i, value := range splitJump {
		splitSpace := strings.Split(value, "\t")
		if i >= 1{
			fmt.Println("ID",splitSpace[3])
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}