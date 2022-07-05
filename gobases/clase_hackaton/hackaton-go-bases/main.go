package main

import (
	//"github.com/bootcamp-go/hackaton-go-bases/internal/interfaces"
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/interfaces"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	// Funcion para obtener tickets del archivo csv
	var booking = service.NewBookings()

	//Leo un registro
	registro, _ := booking.Read(5)
	fmt.Println("Registro leído: ", registro)

	//Creo y agrego un registro
	nuevoRegistro := interfaces.Ticket{Id: 9999,
		Names:       "Cosme Fulanito",
		Email:       "cosme.fulanito@gmail.com",
		Destination: "Miami",
		Date:        "17:55",
		Price:       50,
	}
	nuevoRegistroAgregado, _ := booking.Create(nuevoRegistro)
	fmt.Println("Registro agregado: ", nuevoRegistroAgregado)

	//Edito un registro
	nuevoRegistroAgregado.Names = "Armando Barreda"
	nuevoRegistroAgregado.Email = "armando.barreda@gmail.com"
	nuevoRegistroAgregado.Date = "15:24"
	nuevoRegistroAgregado.Destination = "Mar del Plata"
	nuevoRegistroAgregado.Price = 25

	registroActualizado, _ := booking.Update(nuevoRegistroAgregado)
	fmt.Println("Registro actualizado", registroActualizado)

	//Elimino un registro
	deleteRegistro, _ := booking.Delete(10)

	fmt.Println("Registro Eliminado", deleteRegistro)

}
