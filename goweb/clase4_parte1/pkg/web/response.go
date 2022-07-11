package web

import "strconv"

/*
	Definiremos la estructura de respuestas con los campos:
	- Code: código de respuesta
	- Error: mensaje de error
	- Data: entidad en caso que la respuesta sea correcta
*/
type Response struct {
	Code string			`json:"code"`
	Data interface{}	`json:"data,omitempty"`		 
	Error string		`json:"error,omitempty`
}

/* Implementamos una función que reciba el code, data y error, y nos devuelva
la estructura Respuesta que definimos */
func NewResponse(code int, data interface{}, err string) Response {
	
	/* En caso que el código sea menor a 300, retornamos una repsuesta correcta 
	con el dato y el error vacío */
	if code < 300 {
		// El código lo recibimos como entero, hacemos la conversión a texto con el paquete strconv
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}

	/* En caso contrario, siendo mayor o igual a 300, retornamos una respuesta incorrecta,
	con el dato en nil y el mensaje de error */
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
}