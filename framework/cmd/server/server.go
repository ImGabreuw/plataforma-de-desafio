package main

import (
	"fmt"
	"github.com/ImGabreuw/plataforma-de-desafios/application/repositories"
	"github.com/ImGabreuw/plataforma-de-desafios/domain"
	"github.com/ImGabreuw/plataforma-de-desafios/framework/utils"
	"log"
)

func main() {
	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Felipe",
		Email:    "felipe@gmail.com",
		Password: "123",
	}

	userRepository := repositories.UserRepositoryDB{Db: db}
	result, err := userRepository.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
