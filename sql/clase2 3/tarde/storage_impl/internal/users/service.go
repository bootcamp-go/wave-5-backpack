package users

import (
	"context"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Service interface {
	Store(context.Context, *domains.User) error
	GetOne(context.Context, string) (*domains.User, error)
	Delete(context.Context, string) error
	Update(context.Context, *domains.User) error
}

type service struct {
	rDynamo RepositoryDynamo
}

func NewService(repoDynamo RepositoryDynamo) Service {
	return &service{
		rDynamo: repoDynamo,
	}
}

func (s *service) GetOne(ctx context.Context, id string) (*domains.User, error) {
	return s.rDynamo.GetOne(ctx, id)
}

func (s *service) Store(ctx context.Context, u *domains.User) error {
	return s.rDynamo.Store(ctx, u)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.rDynamo.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, u *domains.User) error {
	user, err := s.rDynamo.GetOne(ctx, u.Id)
	if err != nil {
		return err
	}

	if u.Firstname == "" {
		u.Firstname = user.Firstname
	}

	if u.Lastname == "" {
		u.Lastname = user.Lastname
	}

	if u.Username == "" {
		u.Username = user.Username
	}

	if u.Password == "" {
		u.Password = user.Password
	}

	if u.Email == "" {
		u.Email = user.Email
	}

	if u.IP == "" {
		u.IP = user.IP
	}

	if u.MacAddress == "" {
		u.MacAddress = user.MacAddress
	}

	if u.Website == "" {
		u.Website = user.Website
	}

	if u.Image == "" {
		u.Image = user.Image
	}

	return s.rDynamo.Update(ctx, u)
}
