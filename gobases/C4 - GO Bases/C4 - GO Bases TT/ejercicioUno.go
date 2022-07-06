package main

import(
	"fmt"
	"os"
)

//Leemos el archivo .txt indicado por el usuario
//y capturamos error en caso de haberlo
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
	fmt.Print("Info de txt customers\n--------------------------\n",string(infoCustomers),"\n")
} 

func main(){
	leerArchivo("./customers.txt")
	fmt.Println("\nejecución finalizada")
}