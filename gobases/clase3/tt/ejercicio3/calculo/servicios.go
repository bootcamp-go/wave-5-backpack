package calculo

type Servicios struct {
  Nombre string
  Precio float64
  MinutosTrabajados uint
}

func (s Servicios) calcularPrecio() float64 {
  if s.MinutosTrabajados == 0 {
    return 0
  }

  if s.MinutosTrabajados < 31 {
    return s.Precio
  }

  if s.MinutosTrabajados < 61 {
    return s.Precio * 2
  }

  hs := s.MinutosTrabajados / 60 

  precio := hs * 2

  horas := float64(s.MinutosTrabajados) / 60

  min := horas - float64(hs)

  if 0.5 < min {
    precio += 2
  } else {
    precio ++
  }
  
  final := float64(precio)
  final *= s.Precio

  return final
}

// Recibe un slice de servicios y retorna el total (precio * 30 minutos Trabajados). Por canal
func SumarServicios(servicios []Servicios, ch chan float64) {
  var total float64

  for _, s := range servicios {
    total += s.calcularPrecio()
  }

  ch <- total
}
