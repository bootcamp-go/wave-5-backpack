package service

import "errors"

func checkEmpty(t Ticket) error {
  if (t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "") {
    return errors.New("err: todos los campos son requeridos")
  }

  if t.Price == 0 {
    return errors.New("err: el precio es requerido")
  }

  return nil
}
