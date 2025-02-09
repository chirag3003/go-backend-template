package repository

import "github.com/chirag3003/go-backend-template/db"

var conn db.Connection

type Repository struct {
	User UserRepository
}

func Setup(connection db.Connection) *Repository {
	conn = connection
	return &Repository{
		User: NewUserRepository(),
	}
}
