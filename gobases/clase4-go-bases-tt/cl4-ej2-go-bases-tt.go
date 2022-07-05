/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Registrando clientes
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		The same study of the previous exercise, requests a functionality to
		be able to register data of new clients. The data required to register
		a customer are:
			- File
			- Name and Surname
			- ID NUMBER
			- Telephone number
			- Address
		● Task 1: The docket number must be assigned or generated separately
				  and separately and in a prior to the uploading of the
				  remaining expenses.  Develop and implement a function to
				  generate an ID that you will then use to assign as a value
				  to "Legajo". If for any reason this function returns "nil",
				  it must generate a panic that interrupts the execution and aborts.
		● Task 2: Before registering a client, you must verify if the client
				  already exists. To do this, you need to read the data
				  from a .txt file.  Somewhere in your code, you implement the
				  function to read a file called " customers.txt " (as in the
				  previous exercise, this file does not exist, so the function
				  that tries to read it will return an error).
				 You must properly handle that error as we have seen so far.
				 	That error should :
						1.- generate a panic ;
						2.- launch by console the message: " error: the
						indicated file was not found or is corrupted " ,
						and continue with the execution of the program
						normally, and continue with the execution of the
						program normally .
		● Task 3: After trying to verify whether the client to be
				  registered already exists, develop a function to validate
				  that all data to register for a client contains a non-zero value.
				  This function must return at least two values.  One of the values
				  one of the returned values must be of type error in case a zero
				  value is entered as a parameter (remember the zero values of each
				  data type, e.g. 0, "", nil ).
		● Task 4: Before the end of the execution, even if panics occur, the
				  following messages should be printed on the console by console:
				  "End of execution ", "Several runtime errors were detected
				  runtime errors were detected " and "No files were left open"
				  (in that order). Use defer to meet this requirement.

		General requirements:
			● Use recover to recover the value of panics that may arise
			  (except in. task 1).
			● Remember to perform the necessary validations for each return that
			  may contain an error value (e.g. those attempting to read files).
			  Generate some error, customizing it to your liking, using one of
			  the functions that GO provides for this purpose (also performs that
			  GO provides for it (it also performs the pertinent validation for
			  the case of the error returned)

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

//	STRUCTS
type errores struct {
	err error
}

type cliente struct {
	legajo         string
	nombreApellido string
	DNI            string
	TEL            float64
	domicilio      string
}

//	FUNCTIONS
func leerArchivo(nameFile string) *os.File {
	fileData, err := os.Open(nameFile)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error: el archivo indicado no fue encontrado o está dañado\n", err)
		}
	}()

	if err != nil {
		panic(err)
	}
	fmt.Println("\n\t<<Archivo Cargado>>")
	return fileData
}

func generadorID(c *cliente, nombreApellido string) *cliente {
	if nombreApellido == "" {
		return nil
	}
	number := fmt.Sprint(rand.Intn(100))
	ID := number + nombreApellido
	c.legajo = ID
	return c
}

func verificarCliente(dni string, tel float64, domicilio string) []error {
	var listErr []error
	if dni == "" {
		listErr = append(listErr, errors.New("error. *DNI* no debe ser [vacío]"))
	}
	if tel <= 0 {
		listErr = append(listErr, errors.New("error. *Telefono* debe ser mayor a O [cero]"))
	}
	if domicilio == "" {
		listErr = append(listErr, errors.New("error. *Domicilio* no debe ser [vacío]"))
	}
	if listErr == nil {
		return nil
	}

	return listErr
}

func registroCliente(nombreApellido string, dni string, tel float64, domicilio string) *cliente {
	newClient := &cliente{}

	newClient = generadorID(newClient, nombreApellido)
	if newClient == nil {
		panic("error. El generador ID fallo - Verifique los datos ingresados\n")
	} else {
		fmt.Println("\nFin de la ejecucion")
	}

	err := verificarCliente(dni, tel, domicilio)
	if err != nil {
		fmt.Println("\nSe detectaron varios errores en tiempo de ejecución:")
		for _, i := range err {
			fmt.Println("\t", i)
		}
	} else {
		fmt.Println("> Se ejecuto correctamente\t👍")

		newClient.DNI = dni
		newClient.TEL = tel
		newClient.nombreApellido = nombreApellido

	}

	file := leerArchivo("customers.txt")
	if file != nil {
		fmt.Println(" No han quedado archivos abiertos")
		fmt.Println("\t", file)
		file.Close()
	}

	return newClient
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Impuestos de Salario #2 ||")

	usr1 := registroCliente("Sancho Gacho", "", 0, "")

	if usr1 == nil {
		fmt.Println("\t**Usuario sin registrar**")
	} else {
		fmt.Println("\n", usr1)
	}
}
