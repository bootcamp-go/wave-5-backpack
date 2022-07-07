package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ObtenerUsuarios() []usuario {
	users := []usuario{
		{Id: 1, Nombre: "Leonardo", Apellido: "Da Vinci", Email: "leodavinci@gmail.com", Edad: 504, Altura: 163, Activo: false, FechaDeCreacion: "03/09/1540"},
		{Id: 2, Nombre: "Salvador", Apellido: "Dali", Email: "sdali@gmail.com", Edad: 120, Altura: 172, Activo: false, FechaDeCreacion: "03/09/1940"},
		{Id: 3, Nombre: "Pablo", Apellido: "Picasso", Email: "ppicasso@gmail.com", Edad: 115, Altura: 160, Activo: true, FechaDeCreacion: "03/09/1980"},
	}
	return users
}

func findById(lista []usuario, identificador int) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Id == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}

func findByName(lista []usuario, identificador string) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Nombre == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}
func findBySurname(lista []usuario, identificador string) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Apellido == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}
func findByEmail(lista []usuario, identificador string) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Email == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}
func findByEdad(lista []usuario, identificador int) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Edad == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}
func findByAltura(lista []usuario, identificador int) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Altura == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}

func findByActivo(lista []usuario, identificador bool) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.Activo == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}

func findByFecha(lista []usuario, identificador string) []usuario {
	var listaResultado []usuario
	for _, u := range lista {
		if u.FechaDeCreacion == identificador {
			listaResultado = append(listaResultado, u)
		}
	}
	if len(listaResultado) != 0 {
		return listaResultado
	}
	return []usuario{}
}

type usuario struct {
	Id              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Edad            int    `json:"edad"`
	Altura          int    `json:"altura"`
	Activo          bool   `json:"activo"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}

func HandlerPaginaPrincipal(c *gin.Context) {
	c.String(200, "Bienvenido a la pagina principal!")
}
func HandlerGetAll(c *gin.Context) {
	var users = ObtenerUsuarios()
	fmt.Println(users)
	c.JSON(200, (users))
}

func HandlerFiltrarPorCampos(c *gin.Context) {
	var users = ObtenerUsuarios()
	id := c.Query("id")
	if id != "" {
		i, _ := strconv.Atoi(id)
		users = findById(users, i)
	}
	name := c.Query("nombre")
	if name != "" {
		users = findByName(users, name)
	}
	surname := c.Query("apellido")
	if surname != "" {
		users = findBySurname(users, surname)
	}
	email := c.Query("email")
	if email != "" {
		users = findByEmail(users, email)
	}
	edad := c.Query("edad")
	if edad != "" {
		e, _ := strconv.Atoi(edad)
		users = findByEdad(users, e)
	}
	altura := c.Query("altura")
	if altura != "" {
		a, _ := strconv.Atoi(altura)
		users = findByAltura(users, a)
	}
	activo := c.Query("activo")
	if activo != "" {
		a, _ := strconv.ParseBool(activo)
		users = findByActivo(users, a)
	}
	fecha := c.Query("fecha_de_creacion")
	if fecha != "" {
		users = findByFecha(users, fecha)
	}

	if len(users) > 0 {
		c.JSON(200, gin.H{
			"productos": users,
		})
	} else {
		c.JSON(404, gin.H{
			"error": "No hay usuarios",
		})
	}

}
func main() {

	router := gin.Default()
	router.GET("/", HandlerGetAll)

	grupo := router.Group("/usuarios")
	{
		//http://localhost:8080/usuarios/?nombre=Leonardo
		grupo.GET("/", HandlerFiltrarPorCampos)

	}

	router.Run(":8080")
}
