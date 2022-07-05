package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Customer struct {
	Documents int64
	Name      string
	Last_name string
	DNI       uint64
	Cellphone uint64
	Address   string
}

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	return data, err
}

func GenerateID() (int64, error) {
	return rand.Int63n(1000000), nil
}

func ValidateCustomerData(customer Customer) (bool, error) {
	if customer.DNI == 0 {
		return false, errors.New("DNI no puede ser 0")
	}
	if customer.Name == "" {
		return false, errors.New("Nombre no puede estar vacio")
	}
	if customer.Last_name == "" {
		return false, errors.New("Apellido no puede estar vacio")
	}
	if customer.Documents == 0 {
		return false, errors.New("Legajo no puede ser 0")
	}
	if customer.Address == "" {
		return false, errors.New("Domiclio no puede estar vacio")
	}
	if customer.Cellphone == 0 {
		return false, errors.New("Telefono no puede ser 0")
	}

	return true, nil
}

func searchCustomer([]byte) (bool, error) {
	return true, nil
}

func VerifyCustomer(c Customer) (bool, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	data, err := ReadFile("customers.txt")
	if err != nil {
		return false, err
	}
	searchCustomer(data)
	isValid, validationErr := ValidateCustomerData(c)
	if !isValid {
		return false, validationErr
	}
	return false, nil
}

func createCustomer(c Customer) bool {
	var detectedErr bool
	defer func() {
		fmt.Println("Fin de la ejecucion")
		if detectedErr {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("No han quedado archivos abiertos")
	}()
	id, err := GenerateID()
	if err != nil {
		detectedErr = true
		panic("No se ha asignado un legado valido")
	}
	c.Documents = id
	verified, err := VerifyCustomer(c)
	if err != nil {
		detectedErr = true
		panic(err)
	}
	return verified
}

func main() {
	newCustomer := Customer{
		Name:      "Pepe",
		Last_name: "Perez",
		DNI:       10,
		Cellphone: 305701231,
		Address:   "Cra 98 7-57",
	}

	createCustomer(newCustomer)

}
