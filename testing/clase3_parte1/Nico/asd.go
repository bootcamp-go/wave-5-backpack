/*
func (m *MockRepositorySection) UpdateName(id int, nombre string) (*Product, error) {
	var updated bool = false
	var products []*Product
	if err := m.db.Read(&products); err != nil {
		return nil, err
	}

	var product *Product
	for _, value := range products {
		if value.Id == id {
			value.Nombre = nombre
			product = value
			updated = true
		}
	}

	if !updated {
		return nil, fmt.Errorf("producto id %d no encontrado", id)
	}

	if err := r.db.Write(&products); err != nil {
		return nil, err
	}

	return product, nil
}
*/