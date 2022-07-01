package main

import (
	"fmt"
	"os"
)

// ===========================================
// ================= Structs =================
// ===========================================

// ===========================================
// ================ Funciones ================
// ===========================================

func leerArchivo(archivo string) {
	file, err := os.ReadFile(archivo)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(string(file))
}

// ===========================================
// ================== Main ===================
// ===========================================

func main() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Ejecución finalizada")
	}()

	leerArchivo("customers.txt")
}
