package usuarios

type Usuarios struct {
  Nombre string
  Apellido string
  edad uint8
  correo string
  password string
}

func (u *Usuarios) CambiarNombre(nombre, apellido string) {
  u.Nombre = nombre
  u.Apellido = apellido
}

func (u *Usuarios) CambiarEdad(edad uint8) {
  u.edad = edad
}

func (u *Usuarios) CambiarCorreo(correo string) {
  u.correo = correo
}

func (u *Usuarios) CambiarPassword(password string) {
  u.password = password
}
