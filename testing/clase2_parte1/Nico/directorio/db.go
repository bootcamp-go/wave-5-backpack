package directorio

type DB interface {
	BuscarPorNombre(nombre string) string
	BuscarPorTelefono(telefono string) string
	AgregarEntrada(nombre, telefono string) error
}
