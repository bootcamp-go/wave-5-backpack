package main

import (
	"fmt"
)

// ===========================================
// ================= Structs =================
// ===========================================

type MiError struct{}

func (this *MiError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

// ===========================================
// ================ Funciones ================
// ===========================================

func verificaciones(salary int) error {
	if salary < 150000 {
		return &MiError{}
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
