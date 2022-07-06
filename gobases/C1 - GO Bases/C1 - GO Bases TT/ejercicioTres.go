package main

import "fmt"

func main(){
	//con declaración corta
	switch month := "septiembre"; month{
	case "enero","febrero","marzo","abril","mayo","junio","julio","agosto","septiembre","octubre","noviembre","diciembre":
		fmt.Printf("%s es un mes del año\n", month)
	default:
		fmt.Printf("%s no es un mes del año\n",month)
	}
}