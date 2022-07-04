package calculo

type Productos struct {
  Nombre string
  Precio float64
  Cantidad uint
}

// Recibe un slice de productos y retorna el total (precio * cantidad). Por canal
func SumarProductos(productos []Productos, ch chan float64) {
  var total float64

  for _, v := range productos {
    total += v.Precio * float64(v.Cantidad)
  }
  ch <- total
}
