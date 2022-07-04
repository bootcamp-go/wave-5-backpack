package main

type salary struct {
	Monto   int
	Mensaje string
}

func (s *salary) Error() string {
	return s.Mensaje
}

func impuestoSalary(salarys int) (int, error) {
	if salarys < 150000 {
		return salarys, &salary{
			Monto:   salarys,
			Mensaje: "error: el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return salarys, nil
	}
}
