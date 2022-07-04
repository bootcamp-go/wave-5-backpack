package calculo

type Mantenimiento struct {
  Nombre string
  Precio float64
}

// Recibe un slice de mantenimiento y retorna el precio total. Por canal
func SumarMantenimiento(mantenimientos []Mantenimiento, ch chan float64) {
  var total float64

  for _, v := range mantenimientos {
    total += v.Precio
  }

  ch <- total // Enviando el total al canal
}
