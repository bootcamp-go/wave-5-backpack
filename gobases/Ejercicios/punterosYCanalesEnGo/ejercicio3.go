package main

type servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type mantenimiento struct {
	Nombre string
	Precio float64
}

func newServicio(nombre string, precio float64, minutosTrabajados int) servicio {
	return servicio{
		Nombre:            nombre,
		Precio:            precio,
		MinutosTrabajados: minutosTrabajados,
	}
}

func newMantenimiento(nombre string, precio float64) mantenimiento {
	return mantenimiento{
		Nombre: nombre,
		Precio: precio,
	}
}

func sumarServicios(s []servicio, servi chan float64) {
	total := 0.0

	for _, servicio := range s {
		mediaHsTrabajadas := servicio.MinutosTrabajados / 30
		if servicio.MinutosTrabajados <= 30 {
			total += servicio.Precio
		} else {
			total += servicio.Precio * float64(mediaHsTrabajadas)
		}
	}

	servi <- total
}

func sumarMantenimiento(m []mantenimiento, man chan float64) {
	total := 0.0

	for _, mantenimiento := range m {

		total += mantenimiento.Precio
	}

	man <- total
}
