package repositories

import (
	"github.com/ImGabreuw/plataforma-de-desafios/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (userRepository UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	err = userRepository.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}
