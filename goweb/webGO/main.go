package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id             string `json:"id"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Email          string `json:"email"`
	Edad           int    `json:"edad"`
	Altura         int    `json:"altura"`
	Activo         bool   `json:"activo"`
	Fecha_creacion string `json:"fecha_creacion"`
}

func str_2_int(pal_num string) int {
	num, e := strconv.Atoi(pal_num)
	if e == nil {
		fmt.Printf("%T \n %v", num, num)
	}
	return num
}
func int_2_str(num int64) string {
	return fmt.Sprintf("%d", num)
}

//"github.com/delrio_raul/web-server"
func main() {
	fmt.Println("hola")
	jsonFile, err := os.Open("usuarios.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	fmt.Println(jsonFile)
	var listUsur []usuario
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(byteValue)
	json.Unmarshal(byteValue, &listUsur)
	fmt.Println(listUsur)

	router := gin.Default()
	router.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hola miguel",
		})
	})
	atributos := map[string]int{
		"id":             0,
		"Nombre":         1,
		"Apellido":       2,
		"Email":          3,
		"Edad":           4,
		"Altura":         5,
		"Activo":         6,
		"Fecha_creacion": 7,
	}
	router.GET("/usuarios", func(ctx *gin.Context) {
		params := ctx.Request.URL.Query()
		fmt.Println(params)
		fmt.Println("hola2")

		if len(params) != 0 {
			var listEspecial []usuario
			var v reflect.Value
			var t reflect.Type
			var contador int
			var atribStri string
			keys := make([]string, 0, len(params))
			for p := range params {
				keys = append(keys, p)
			}
			for _, usuario := range listUsur {
				contador = 0
				t = reflect.TypeOf(usuario)
				v = reflect.ValueOf(usuario)
				for _, k := range keys {
					tipoAtrib := t.Field(atributos[k]).Type.Kind().String()
					if tipoAtrib == "int" {
						atribStri = int_2_str(v.Field(atributos[k]).Int())
					} else if tipoAtrib == "bool" {
						atribStri = strconv.FormatBool(v.Field(atributos[k]).Bool())
					} else {
						atribStri = v.Field(atributos[k]).String()
					}
					fmt.Println(tipoAtrib)
					fmt.Println(params.Get(k))
					fmt.Println(atribStri)
					fmt.Println(v.Field(atributos[k]).String())
					if params.Get(k) == atribStri {
						contador++
					}
				}
				if contador == len(params) {
					listEspecial = append(listEspecial, usuario)
				}
			}

			// listFiltros := filtros(params)
			ctx.JSON(200, listEspecial)

		} else {
			fmt.Println("entre carajo")
			ctx.JSON(200, listUsur)
		}

	})
	router.GET("/usuarios/:id", func(ctx *gin.Context) {
		block := true
		for _, user := range listUsur {
			if user.Id == ctx.Param("id") {
				ctx.JSON(200, user)
				block = false
			}
		}
		if block {
			ctx.JSON(404, "no existe vuelva pronto")
		}
	})

	router.Run()
}

// func filtros(listUsur []usuario,params url.Values) string {
// 	e := reflect.ValueOf(&book).Elem()
// 	for _,usur:= range listUsur{
// 		for _,atrib := range reflect.ValueOf(usur){

// 		}

// 	}
// }
