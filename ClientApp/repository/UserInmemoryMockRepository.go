package repository

import "github.com/virtouso/WhatsappClientServer/ClientApp/model/domain"

type InMemoryMockRepository struct {
}

func (i InMemoryMockRepository) ReadAll() []*domain.User {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryMockRepository) Init() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryMockRepository) Create(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryMockRepository) Read(id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryMockRepository) Update(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i InMemoryMockRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

// simple way to force implement interface. no time to implement mocking
func makeInstance() {
	var _ IUserRepository = InMemoryMockRepository{}
}
