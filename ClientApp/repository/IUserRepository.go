package repository

import "github.com/virtouso/WhatsappClientServer/ClientApp/model/domain"

var UserRepo IUserRepository

type IUserRepository interface {
	Init() (bool, error)
	Create(user *domain.User) (*domain.User, error)
	Read(id string) (*domain.User, error)
	ReadAll() []*domain.User
	Update(user *domain.User) (*domain.User, error)
	Delete(id string) error
}
