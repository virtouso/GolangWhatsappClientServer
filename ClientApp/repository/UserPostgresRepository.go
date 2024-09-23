package repository

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/virtouso/WhatsappClientServer/ClientApp/model/domain"
	"github.com/virtouso/WhatsappClientServer/ClientApp/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresUserRepository struct{}

func (m PostgresUserRepository) ReadAll() []*domain.User {
	var users []*domain.User
	_ = Db.Find(&users)
	return users
}

func (m PostgresUserRepository) Create(user *domain.User) (*domain.User, error) {
	err := Db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m PostgresUserRepository) Read(id string) (*domain.User, error) {
	user := domain.User{}
	Db.Where("id=?", id).First(&user)

	if user == (domain.User{}) {
		return nil, errors.New(("not found"))
	}
	return &user, nil
}

func (m PostgresUserRepository) Update(user *domain.User) (*domain.User, error) {
	result := Db.Where("id=?", user.ID).Updates(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (m PostgresUserRepository) Delete(id string) error {
	var wallet = domain.User{}
	result := Db.Where("id = ?", id).Delete(wallet)
	return result.Error
}

var Db *gorm.DB

func (m PostgresUserRepository) Init() (bool, error) {
	var err error
	dsn := os.Getenv(shared.UserDbConKey)
	fmt.Println("dsn : ", dsn)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Migrate()
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return false, err
	}

	return true, nil
}

func Migrate() {
	Db.AutoMigrate(&domain.User{})
	log.Println("Database Migration Completed...")
}

func makePostgresInstance() {
	var _ IUserRepository = PostgresUserRepository{}
}
