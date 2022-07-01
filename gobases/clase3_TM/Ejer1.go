package main

import ("fmt"
		"os")


type products struct{
	id int
	price float64
	quantity int
}
func main(){
	fileName := "./myFile.csv"
	var p1 products
	// para ingresar datos de productos ingrese add como texto seguido de enter
	for{
		var action string
		fmt.Println("ingrese add para agregar parametros, read para leer el archivo y exit para terminar")
		fmt.Scanln(&action)

		switch action {
		case "add":
			add(&p1)
			appendCsvline(p1,fileName)
		case "read":
			fmt.Println("no definido")
		case "exit":
			os.Exit(0)	
		default:
			fmt.Println("valor no correcto")
		}
	}
}

func add(p *products){
	fmt.Println("ingrese el Id del producto")
	fmt.Scanln(&p.id)
	fmt.Println("ingrese el precio del producto")
	fmt.Scanln(&p.price)
	fmt.Println("ingrese el cantidad del producto")
	fmt.Scanln(&p.quantity)
}

func appendCsvline(p products, fileName string){


	// generamos el formato de una linea para el CSV
	format :=fmt.Sprintf("%d,%.2f,%d",p.id,p.price,p.quantity)
	// casteamos a un slice de tipo byte, por si no existe archivoad
	 data := []byte(format)

	_,err:=os.Stat(fileName)

	if err != nil {
		fmt.Println(err)
		err =os.WriteFile(fileName,data,0644)
		if err!= nil{
			fmt.Printf("Error escritura: %v", err)
		}
	}else{
		fmt.Println("si existe, agregando nueva linea")
		file,_ := os.ReadFile(fileName)
		format=string(file)+"\n"+format
		data=[]byte(format)
		err =os.WriteFile(fileName,data,0644)

	}

}