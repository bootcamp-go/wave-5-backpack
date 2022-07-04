package clientes

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
  Legajo uint64
  Nombre string
  Apellido string
  DNI uint
  Domicilio string
  NumeroTel string 
}

// Retorna un id para usar en el legajo
func NuevoId() uint64 {
  return rand.Uint64()
}

// Check si el id ya existe
func CheckId(id uint64) error {
  defer func () {
    err := recover()

		if err != nil {
			log.Printf("panic detectado: %v\n", err)
		}

		log.Println("No han quedado archivos abiertos")
  }()

  data, err := os.ReadFile("./customers.txt")
  if err != nil {
    panic("error: el archivo indicado no fue encontrado o está dañado")
  }

  // Leer data
  lineas := strings.Split(string(data), "\n")

  for _, v := range lineas {
  	idLeido, err := strconv.ParseUint(v, 10, 64)
  	if err != nil {
  		return err
  	}

  	if id == idLeido {
  		return errors.New("ya existe cliente con ese id")
  	}
  }

  return nil
}

func NuevoCliente(legajo uint64, dni uint, nombre, apellido, domicilio, numTel string) (*Cliente, error) {
	defer func() {
		err := recover()	

		if err != nil {
			log.Printf("err: ocurrieron errores durante la validación")
		}
	}()

	if legajo == 0 {
		panic("panic: num de legajo es 0")
	}

	if dni == 0 {
		panic("panic: DNI no puede ser 0")
	}

	if nombre == "" {
		panic("panic: nombre es requerido")
	}
	
	if apellido == "" {
		panic("panic: apellido es requerido")
	}

	if domicilio == "" {
		panic("panic: domicilio es requerido")
	}

	if numTel == "" {
		panic("panic: numTel es requerido")
	}

	return &Cliente{
		Legajo: legajo,
		Nombre: nombre,
		Apellido: apellido,
		Domicilio: domicilio,
		NumeroTel: numTel,
		DNI: dni,
	},nil
}
