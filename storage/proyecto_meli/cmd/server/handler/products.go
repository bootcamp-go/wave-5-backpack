package handler

import (
	"fmt"
	"proyecto_meli/internal/products"
	"proyecto_meli/pkg/web"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creacion"`
}

const (
	BAD_TOKEN         = "Token inválido"
	BAD_ID            = "ID inválido"
	PRODUCT_NOT_FOUND = "Producto no encontrado"
	F_REQUIRED_ITEM   = "El campo %s es requerido."
	F_REQUIRED_ITEMS  = "Los campos %s son requeridos."
	F_BAD_ITEM        = "El campo %s es invalido."
	F_BAD_ITEMS       = "Los campos %s son invalidos."
	S_ID              = "id"
	S_NAME            = "nombre"
	S_COLOR           = "color"
	S_PRICE           = "precio"
	S_STOCK           = "stock"
	S_CODE            = "codigo"
	S_PUBLISH         = "publicado"
	S_DATE            = "fecha_de_creacion"
	PRODUCT_DELETE_ID = "El producto %d ha sido eliminado"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.service.GetAll(c)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, products, ""))
	}
}

// ProductById godoc
// @Summary Product by ID
// @Tags Products
// @Description get product by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/:id [get]
func (p *Product) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(500, web.NewResponse(500, nil, BAD_ID))
			return
		}
		product, err := p.service.GetById(c, id)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, PRODUCT_NOT_FOUND))
			return
		}
		c.JSON(200, web.NewResponse(200, product, ""))
	}
}

// ListFilterProducts godoc
// @Summary List products by filter
// @Tags Products
// @Description get products by filter
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Query product body
// @Success 200 {object} web.Response
// @Router /products/filter [get]
func (p *Product) FilterList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errorValidacion []string
		var name, color, codigo string
		var stock int = -1
		var id int
		var price float64
		var publicado bool = true
		var fecha string
		idQ := c.Query("id")
		if idQ != "" {
			Tem, err := strconv.Atoi(idQ)
			if err != nil {
				errorValidacion = append(errorValidacion, S_ID)
			} else {
				id = Tem
			}
		}
		nameQ := c.Query("nombre")
		if nameQ != "" {
			name = nameQ
		}
		colorQ := c.Query("color")
		if colorQ != "" {
			color = colorQ
		}
		priceQ := c.Query("precio")
		if priceQ != "" {
			Tem, err := strconv.ParseFloat(priceQ, 64)
			if err != nil {
				errorValidacion = append(errorValidacion, S_PRICE)
			} else {
				price = Tem
			}
		}
		stockQ := c.Query("stock")
		if stockQ != "" {
			Tem, err := strconv.Atoi(stockQ)
			if err != nil {
				errorValidacion = append(errorValidacion, S_STOCK)
			} else {
				stock = Tem
			}
		}
		codigoQ := c.Query("codigo")
		if codigoQ != "" {
			codigo = codigoQ
		}
		publicadoQ := c.Query("publicado")
		if publicadoQ != "" {
			Tem, err := strconv.ParseBool(publicadoQ)
			if err != nil {
				errorValidacion = append(errorValidacion, S_PUBLISH)
			}
			publicado = Tem
		}
		fechaQ := c.Query("fecha")
		if fechaQ != "" {
			_, err := time.Parse("2006-01-02", fechaQ)
			if err != nil {
				fmt.Println(err)
				errorValidacion = append(errorValidacion, S_DATE)

			} else {
				fecha = fechaQ
			}
		}

		products, err := p.service.FilterList(c, id, name, color, price, stock, codigo, publicado, fecha)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if len(errorValidacion) > 0 {
			if len(errorValidacion) > 1 {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_BAD_ITEMS, strings.Join(errorValidacion, ", "))))
				return
			} else {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_BAD_ITEM, errorValidacion[0])))
				return
			}
		}

		if len(products) > 0 {
			c.JSON(200, web.NewResponse(200, products, ""))
			return
		}
		c.JSON(404, web.NewResponse(404, nil, PRODUCT_NOT_FOUND))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(500, err.Error())
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Color == "" {
			errStruct = append(errStruct, "color")
		}
		if request.Precio < 0 {
			errStruct = append(errStruct, "precio")
		}
		if request.Stock < 0 {
			errStruct = append(errStruct, "stock")
		}
		if request.Codigo == "" {
			errStruct = append(errStruct, "codigo")
		}
		if request.FechaCreacion == "" {
			errStruct = append(errStruct, "fecha_de_creacion")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, strings.Join(errStruct, ", "))))
				return
			} else {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, errStruct[0])))
				return
			}
		}
		product, err := p.service.Store(c, request.Nombre, request.Color, request.Precio, request.Stock, request.Codigo, request.Publicado, request.FechaCreacion)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, product, ""))
	}
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Products
// @Description Update produc by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/:id [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, BAD_ID))
			return
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Color == "" {
			errStruct = append(errStruct, "color")
		}
		if request.Precio < 0 {
			errStruct = append(errStruct, "precio")
		}
		if request.Stock < 0 {
			errStruct = append(errStruct, "stock")
		}
		if request.Codigo == "" {
			errStruct = append(errStruct, "codigo")
		}
		if request.FechaCreacion == "" {
			errStruct = append(errStruct, "fecha_de_creacion")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, strings.Join(errStruct, ", "))))
				return
			} else {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, errStruct[0])))
				return
			}
		}
		product, err := p.service.Update(c, id, request.Nombre, request.Color, request.Precio, request.Stock, request.Codigo, request.Publicado, request.FechaCreacion)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, product, ""))
	}
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description Delete produc by ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/:id [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, BAD_ID))
			return
		}
		err = p.service.Delete(c, id)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, fmt.Sprintf(PRODUCT_DELETE_ID, id), ""))
	}
}

// UpdateParcialProduct godoc
// @Summary Update parcial product
// @Tags Products
// @Description Update produc by ID with Name and Preci
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/:id [patch]
func (p *Product) Update_Name_Price() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Print(err)
			c.JSON(400, web.NewResponse(400, nil, BAD_ID))
			return
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var errStruct []string
		if request.Nombre == "" {
			errStruct = append(errStruct, "nombre")
		}
		if request.Precio < 0 {
			errStruct = append(errStruct, "precio")
		}
		if len(errStruct) > 0 {
			if len(errStruct) > 1 {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, strings.Join(errStruct, ", "))))
				return
			} else {
				c.JSON(500, web.NewResponse(500, nil, fmt.Sprintf(F_REQUIRED_ITEMS, errStruct[0])))
				return
			}
		}
		product, err := p.service.Update_Name_Price(c, id, request.Nombre, request.Precio)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, product, ""))
	}
}

func (p *Product) GetProductByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("nombre")

		products, err := p.service.GetProductByName(c, name)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "producto no encontrado"))
			return
		}
		c.JSON(200, web.NewResponse(200, products, ""))
	}
}
