package middlewares

import "github.com/chirag3003/go-backend-template/repository"

var repo repository.Repository

type Middlewares struct{}

func Setup(repository repository.Repository) *Middlewares {
	repo = repository
	return &Middlewares{}
}
