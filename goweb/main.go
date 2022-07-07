package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Estructura de usuarios
type Usuarios struct {
	Id            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Edad          int       `json:"edad"`
	Altura        float64   `json:"altura"`
	Activo        bool      `json:"activo"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}

func ObtenerUsuarios() ([]Usuarios, error) {
	// Abrimos el archivo
	jsonFile, err := os.Open("./usuarios.json")
	if err != nil {
		return nil, errors.New("no se puede abrir el archivo")
	}

	// Cerramos el archivo
	defer jsonFile.Close()

	// Creamos el objeto decodificador
	decoder := json.NewDecoder(jsonFile)

	// Creamos el Slice de usuarios
	var users []Usuarios

	// Decodificamos el archivo y se asigna al slice
	err = decoder.Decode(&users)
	if err == io.EOF {
		return nil, errors.New("eof")
	}
	if err != nil {
		return nil, errors.New("error al parsear el archivo a json")
	}

	return users, nil
}

func GetAll(c *gin.Context) {
	// Se obtienen los usuarios
	users, err := ObtenerUsuarios()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		// Devolvemos el JSON al cliente
		c.JSON(200, users)
	}
}

func SearchUser(c *gin.Context) {
	// Se obtienen los usuarios
	users, err := ObtenerUsuarios()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		var filtro []Usuarios
		var temporal []Usuarios

		// Se obtienen los query params
		nombreParam := strings.ToUpper(c.Query("nombre"))
		apellidoParam := strings.ToUpper(c.Query("apellido"))
		emailParam := strings.ToUpper(c.Query("email"))
		edadParam, edadErr := strconv.Atoi(c.Query("edad"))
		alturaParam, alturaErr := strconv.ParseFloat(c.Query("altura"), 64)
		activoParam, activoErr := strconv.ParseBool(c.Query("activo"))
		fechaCreacionParam, fechaErr := time.Parse("2006-01-02", c.Query("fecha_creacion"))

		// Buscamos por nombre
		if nombreParam != "" {
			for _, u := range users {
				if strings.Contains(strings.ToUpper(u.Nombre), nombreParam) {
					filtro = append(filtro, u)
				}
			}
		}

		// Buscamos por apellido
		if apellidoParam != "" {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if strings.Contains(strings.ToUpper(u.Apellido), apellidoParam) {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if strings.Contains(strings.ToUpper(u.Apellido), apellidoParam) {
						filtro = append(filtro, u)
					}
				}
			}
		}

		// Buscamos por Email
		if apellidoParam != "" {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if strings.Contains(strings.ToUpper(u.Email), emailParam) {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if strings.Contains(strings.ToUpper(u.Email), emailParam) {
						filtro = append(filtro, u)
					}
				}
			}
		}

		// Buscamos por edad
		if edadErr == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if edadParam == u.Edad {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if edadParam == u.Edad {
						filtro = append(filtro, u)
					}
				}
			}
		}

		// Buscamos por altura
		if alturaErr == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if alturaParam == u.Altura {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if alturaParam == u.Altura {
						filtro = append(filtro, u)
					}
				}
			}
		}

		// Buscamos por activo
		if activoErr == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if activoParam == u.Activo {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if activoParam == u.Activo {
						filtro = append(filtro, u)
					}
				}
			}
		}

		// Buscamos por fecha
		if fechaErr == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if fechaCreacionParam == u.FechaCreacion {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if fechaCreacionParam == u.FechaCreacion {
						filtro = append(filtro, u)
					}
				}
			}
		}

		if len(filtro) > 0 {
			c.JSON(http.StatusOK, filtro)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	}
}

func UserByID(c *gin.Context) {
	// Se obtienen los usuarios
	users, err := ObtenerUsuarios()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		var filtro []Usuarios
		found := false

		// Se obtiene el parametro ID
		idInt, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Si el ID no tiene el formato correcto, se devuelve 404
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			// Se busca el usuario por ID
			for _, u := range users {
				if u.Id == idInt {
					// Usuario encontrado
					filtro = append(filtro, u)
					found = true
					break
				}
			}

			if found {
				// Se encontró el usuario
				c.JSON(http.StatusOK, filtro)
			} else {
				// No se encontró el usuario
				c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			}
		}
	}
}

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET "/hello"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Arturo"})
	})

	// Devuelve los elementos de los usuarios.json
	router.GET("/usuarios", GetAll)

	// Router para buscar usuarios
	router.GET("/searchUser", SearchUser)

	// Router para buscar usuarios por ID
	router.GET("/usuarios/:id", UserByID)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
