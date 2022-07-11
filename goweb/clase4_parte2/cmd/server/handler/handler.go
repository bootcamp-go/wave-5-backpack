package handler

import (
	"strconv"

	"clase4_parte2/internal/products"
	"clase4_parte2/pkg/store/web"

	"github.com/gin-gonic/gin"
	/*
		"clase4_parte2/docs"
		"github.com/swaggo/files"
		ginSwagger "github.com/swaggo/gin-swagger"
	*/)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

// Update godoc
// @Summary Update products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			token := c.GetHeader("token")
			if token != os.Getenv("TOKEN") {
				c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token invalido"))
				return
			}
		*/
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			//c.JSON(400, gin.H{"error": "Id inválido"})
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"error": err.Error()})
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			//c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			//c.JSON(400, gin.H{"error": "El tipo de producto es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El tipo de producto es requerido"))
			return
		}

		if req.Count == 0 {
			//c.JSON(400, gin.H{"error": "La cantidad es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}

		if req.Price == 0 {
			//c.JSON(400, gin.H{"error": "El precio es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		p, err := p.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//c.JSON(200, p)
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Actualiza nombre de productos godoc
// @Summary Actualiza products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [PATCH]
func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token invalido"))
			return
		}*/

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			//c.JSON(400, gin.H{"error": "Id inválido"})
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"error": err.Error()})
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			//c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		p, err := p.service.UpdateName(int(id), req.Name)
		if err != nil {
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		//c.JSON(200, p)
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Elimina productos godoc
// @Summary Elimina products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [DELETE]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token invalido"))
			return
		}*/

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			//c.JSON(400, gin.H{"error": "Id inválido"})
			c.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		//c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Lista productos godoc
// @Summary Lista products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			token := c.GetHeader("token")
			if token != os.Getenv("TOKEN") {
				//c.JSON(401, gin.H{"error": "Token inválido"})
				c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token invalido"))
				return
			}
		*/
		p, err := p.service.GetAll()
		if err != nil {
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		//c.JSON(200, p)
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token invalido"))
			return
		}*/

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			//c.JSON(400, gin.H{"error": err.Error()})
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			//c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			//c.JSON(400, gin.H{"error": "El tipo de producto es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El tipo de producto es requerido"))
			return
		}

		if req.Count == 0 {
			//c.JSON(400, gin.H{"error": "La cantidad es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "La cantidad es requerido"))
			return
		}

		if req.Price == 0 {
			//c.JSON(400, gin.H{"error": "El precio es requerido"})
			c.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		p, err := p.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		//c.JSON(200, gin.H{"products": p})
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}
