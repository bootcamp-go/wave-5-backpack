package directorio

import "fmt"

type engine struct {
	database DB
}

func NewEngine(db DB) engine {
	return engine{database: db}
}

func (e engine) GetVersion() string {
	return "1.0"
}

func (e engine) FindByName(name string) string {
	if len(name) > 3 {
		return e.database.BuscarPorNombre(name)
	}
	return ""
}

func (e engine) FindByTelephone(telephone string) string {
	if len(telephone) > 5 {
		return e.database.BuscarPorTelefono(telephone)
	}
	return ""
}

func (e engine) AddEntry(name, telephone string) error {
	if len(name) > 3 && len(telephone) > 5 {
		return e.database.AgregarEntrada(name, telephone)
	}
	return fmt.Errorf("name %s and telephone %s cannot be added to the registry", name, telephone)
}
