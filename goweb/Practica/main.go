package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	FIELD_EMPTY = "el campo %s no puede estar vacio"
)

type Product struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      uint64  `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
}

type ProductRequest struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      uint64  `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
}

var prList []Product
var IdGeneral int

func getProductList() {
	data, err := os.ReadFile("products.json")
	if err != nil {
		panic(err)
	}
	var res []Product
	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	IdGeneral = len(prList)
	prList = res
}

func getProductById(productList []Product, identifier uint) (Product, error) {
	for _, product := range productList {
		if product.Id == identifier {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("cant find product: %d", identifier)
}

func filterList(prList *[]Product, params url.Values) error {
	res := []Product{}
	name := params.Get("name")
	color := params.Get("color")
	price := params.Get("price")
	stock := params.Get("stock")
	code := params.Get("code")
	published := params.Get("published")
	created := params.Get("created")
	for _, product := range *prList {
		valid := true
		if name != "" {
			if product.Name != name {
				valid = false
			}
		}
		if color != "" && valid {
			if product.Color != color {
				valid = false
			}
		}
		if price != "" && valid {
			comparePrice, err := strconv.ParseFloat(price, 64)
			if err != nil {
				return err
			}
			if product.Price != comparePrice {
				valid = false
			}
		}
		if stock != "" && valid {
			compareStock, err := strconv.ParseUint(stock, 10, 64)
			if err != nil {
				return err
			}
			if product.Stock != compareStock {
				valid = false
			}
		}
		if code != "" && valid {
			if product.Code != code {
				valid = false
			}
		}
		if published != "" && valid {
			comparePublished, err := strconv.ParseBool(published)
			if err != nil {
				return err
			}
			if product.Published != comparePublished {
				valid = false
			}
		}
		if created != "" && valid {
			if product.Created_at != created {
				valid = false
			}
		}

		if valid {
			res = append(res, product)
		}
	}
	*prList = res
	return nil
}

func getAll(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	if len(params) > 0 {
		err := filterList(&prList, params)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	ctx.JSON(200, prList)
}

func getById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	product, err := getProductById(prList, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func create(ctx *gin.Context) {
	getProductList()
	var nwRegistro ProductRequest
	ctx.ShouldBindJSON(&nwRegistro)
	errors := validate(nwRegistro)
	if len(errors) > 0 {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}
	/*token*/
	headerToken := ctx.GetHeader("token")
	if headerToken != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "No tiene permisos para realizar la petición solicitada.",
		})
		return
	}
	IdGeneral++
	nwUser := Product{
		Id:         nwRegistro.Id,
		Name:       nwRegistro.Name,
		Color:      nwRegistro.Color,
		Price:      nwRegistro.Price,
		Stock:      nwRegistro.Stock,
		Code:       nwRegistro.Code,
		Published:  nwRegistro.Published,
		Created_at: time.Now().Format("2022-12-25"),
	}

	prList = append(prList, nwUser)
	ctx.JSON(http.StatusOK, nwUser)
}

func validate(user ProductRequest) []string {
	var errors []string
	if user.Name == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "nombre").Error())
	}

	if user.Color == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "apellido").Error())
	}

	if user.Price == 0 {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "email").Error())
	}

	if user.Stock == 0 {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "edad").Error())
	}

	if user.Code == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "estatura").Error())
	}
	return errors
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Request.URL.Query().Get("name")
		if name == "" {
			name = "Anonimo"
		}
		ctx.JSON(200, gin.H{
			"message": "Saludos " + name,
		})
	})

	router.GET("/products", getAll)
	router.GET("/products/:id", getById)
	router.POST("/products", create)
	router.Run()
}
