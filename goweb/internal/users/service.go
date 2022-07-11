package users

import (
	"errors"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	ValidateToken(token string) error
	GetById(id int) (domain.User, error)
	UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	/* 	ValidateReq(domain.User)([]string)
	 */
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

/* type request struct {
	Name       string  `json: "name"`
	Lastname   string  `json: "lastname"`
	Email      string  `json: "email"`
	Age        int     `json: "age"`
	Height     float32 `json: "height"`
	Active     bool    `json: "active"`
	DoCreation string  `json: "doCreation"`
}

func (s *service) ValidateReq(req request) []string {
	var errMsg []string

	if req.Name == "" {
		errMsg = append(errMsg, "Name required")
	}

	if req.Lastname == "" {
		errMsg = append(errMsg, "Lastname required")
	}
	if req.Email == "" {
		errMsg = append(errMsg, "Email required")
	}
	if req.Age == 0 {
		errMsg = append(errMsg, "Age required")
	}
	if req.Height == 0 {
		errMsg = append(errMsg, "Height required")
	}
	if req.DoCreation == "" {
		errMsg = append(errMsg, "Date of creation required")
	}

	return errMsg

} */

func (s *service) ValidateToken(token string) error {
	if token != os.Getenv("TOKEN") {
		err := errors.New("ERROR: Invalid token")
		return err
	}
	return nil
}
func (s *service) GetAll() ([]domain.User, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *service) GetById(id int) (domain.User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	lastID, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastID++

	newUser, err := s.repository.StoreUser(lastID, name, lastname, email, age, height, active, doCreation)

	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}
func (s *service) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	updatedUser, err := s.repository.UpdateUser(id, name, lastname, email, age, height, active, doCreation)

	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}
