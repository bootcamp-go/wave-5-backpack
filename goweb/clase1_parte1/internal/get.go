package internal

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {

	// Filtrando con todos los campos
	var struct_producto producto

	if err := ctx.ShouldBindQuery(&struct_producto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapKeys := make(map[string]interface{})
	mapKeys["string"] = ""
	mapKeys["int"] = 0
	mapKeys["bool"] = false

	mapProduct := jsonToMap(struct_producto)[0]
	var keysList = []string{}
	for key, value := range mapProduct {
		if value != mapKeys[reflect.TypeOf(value).String()] {
			keysList = append(keysList, key)
		}
	}

	productos := jsonToMap(getJson()...)
	var filtered_products []map[string]interface{}
	var filtered_products_empty []map[string]interface{}

	for _, key := range keysList {
		for _, p := range productos {
			if p[key] == mapProduct[key] {
				filtered_products = append(filtered_products, p)
			}
		}
		fmt.Println(filtered_products)
		productos = filtered_products
		filtered_products = filtered_products_empty
	}

	ctx.JSON(200, productos)

	// Filtrando solo por el ID
	/*
		productos := getJson()
		var list []producto
		for _, p := range productos {
			if p.Id == struct_producto.Id {
				list = append(list, p)
			}
		}
		ctx.JSON(200, list) */

}

func jsonToMap(productos ...producto) []map[string]interface{} {
	productoMap := make(map[string]interface{})
	var productoMaps []map[string]interface{}

	for _, p := range productos {
		productoMap["id"] = p.Id
		productoMap["nombre"] = p.Nombre
		productoMap["color"] = p.Color
		productoMap["precio"] = p.Precio
		productoMap["stock"] = p.Stock
		productoMap["codigo"] = p.Codigo
		productoMap["publicado"] = p.Publicado
		productoMap["fecha"] = p.Fecha

		productoMaps = append(productoMaps, productoMap)
		productoMap = make(map[string]interface{})
	}

	return productoMaps
}

func GetById(ctx *gin.Context) {
	productos := getJson()
	id, _ := strconv.Atoi(ctx.Param("id"))
	var producto_id producto
	for _, p := range productos {
		if id == p.Id {
			producto_id = p
			break
		}
	}
	ctx.JSON(200, producto_id)
}
