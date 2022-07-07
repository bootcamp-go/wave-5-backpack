package main

import (
	"errors"
	"fmt"
)

// ===========================================
// ================= Structs =================
// ===========================================

// ===========================================
// ================ Funciones ================
// ===========================================

func verificaciones(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

// ===========================================
// ================== Main ===================
// ===========================================

func main() {
	var salary int = 10000000
	err := verificaciones(salary)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary = 15000
	err = verificaciones(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
