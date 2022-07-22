package handler

import (
	"fmt"
	"goweb/internal/products"
	"goweb/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id            int     `json:"-"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"código"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creación"`
}

type Product struct {
	service products.Service
}

func InitProduct(p products.Service) *Product {
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
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		} */
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// CreateProducts godoc
// @Summary Create products
// @Tags Products
// @Description create products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/create [post]
func (c *Product) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if v := validador(req); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		p, err := c.service.CreateProduct(req.Id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusOK, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		} */
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if v := validador(req); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}
		/* if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es requerido",
			})
			return
		}
		if req.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El color del producto es requerido",
			})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es requerido",
			})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El stock del producto es requerido",
			})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El código del producto es requerido",
			})
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "La fecha de creación del producto es requerida",
			})
			return
		} */
		p, err := c.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}

}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		} */
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id inválido"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK,
			web.NewResponse(http.StatusOK, nil, fmt.Sprintf("El producto %d ha sido eliminado", id)))
	}

}

func (c *Product) UpdateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		} */
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "id inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		p, err := c.service.UpdateOne(int(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}

}

func validador(req request) string {
	var response string
	if req.Nombre == "" {
		response += "Falta el campo nombre "
	}
	if req.Color == "" {
		response += "Falta el campo color "
	}
	if req.Precio == 0 {
		response += "Falta el campo precio "
	}
	if req.Stock == 0 {
		response += "Falta el campo stock "
	}
	if req.Codigo == "" {
		response += "Falta el campo código "
	}
	if req.FechaCreacion == "" {
		response += "Falta el campo fecha de creación "
	}
	return response
}
