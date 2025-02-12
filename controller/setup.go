package controller

import (
	"github.com/chirag3003/go-backend-template/repository"
)

var repo *repository.Repository

type Controllers struct{
  Auth AuthController
  User UserController
}

func Setup(repository *repository.Repository) *Controllers {
	repo = repository
	return &Controllers{
    Auth: newAuthController(),
    User: newUserController(),
  }
}
