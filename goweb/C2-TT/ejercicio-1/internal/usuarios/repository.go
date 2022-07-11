package usuarios

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fecha_creacion" binding:"required"`
}

type Repository interface {
	GetAll() ([]*Usuario, error)
	Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
}

var usuarios []*Usuario
var lastID int

func NewRepository() Repository {
	return &Usuario{}
}

func (u *Usuario) GetAll() ([]*Usuario, error) {
	return usuarios, nil
}

func (u *Usuario) Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	lastID++

	usuario := Usuario{
		ID:            lastID,
		Nombre:        nombre,
		Apellido:      apellido,
		Email:         email,
		Edad:          edad,
		Altura:        altura,
		Activo:        activo,
		FechaCreacion: fecha_creacion,
	}

	usuarios = append(usuarios, &usuario)

	return &usuario, nil
}
