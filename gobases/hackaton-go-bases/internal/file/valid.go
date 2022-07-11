package file

// Retorna true si algun campo es vacio
func checkEmpty(slice []string) bool {
	for _, v := range slice {
		if v == "" {
			return true
		}
	}

	return false
}
