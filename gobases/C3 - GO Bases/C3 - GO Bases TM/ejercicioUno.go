package main

import(
	"fmt"
	"os"
)

func main(){
	infoProductos := [][]string{
		{"ID","Precio","Cantidad"},
		{"1","1200","3"},
		{"2","4000","4"},
		{"3","4200","2"},
	}

	//Creación .csv de productos
	_, err := os.Create("infoProducto.csv")
	if(err!=nil){
		fmt.Println("Problemas con la creación del archivo")
	}

	//Adicioné la opción de añadir lo de otro csv
	data, _ := os.ReadFile("./addProducts.csv")
	fmt.Printf("Info de csv a adicionar\n %v\n",string(data))
	
	//para concatenar los valores a retornar en el csv
	returnValue := ""
	for i := 0; i < len(infoProductos) ; i++ {
		for j := 0; j < 3 ; j++{ 
			//Para mejor visualización
			returnValue += fmt.Sprint(infoProductos[i][j],",\t")
		}
		//Para la visualización en el .csv
		returnValue += fmt.Sprint("\n")
	}

	returnValue += string(data)
	//WriteFile para escribirlo en el csv creado
	dataProducto := []byte(returnValue)
	err = os.WriteFile("./infoProducto.csv", dataProducto, 0644)
}