package main

import ("fmt"
	 	"os")
type alumnos struct{
	nombre string
	apellido string
	DNI 	int
	fecha	string
}

func main(){
	a1:=alumnos{"Elkin","Suarez",105876,"16/04/1994"}

	a1.detalle()

	files,err := os.ReadDir(".")
	if err==nil{
		fmt.Printf("files:%v \n", files)
	}

	text := []byte("clase 3 archivos")
	err =os.WriteFile("./myFile.txt",text,0644)

	if err!= nil{
		fmt.Printf("Error escritura: %v", err)
	}
}


func (a alumnos) detalle(){
	fmt.Printf("Nombre: %s \n",a.nombre)
	fmt.Printf("Nombre: %s \n",a.apellido)
	fmt.Printf("Nombre: %d \n",a.DNI)
	fmt.Printf("Nombre: %s \n",a.fecha)
}
