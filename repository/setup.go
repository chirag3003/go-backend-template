package repository

import "github.com/chirag3003/go-backend-template/db"

var conn db.Connection

type Repository struct {
	User  UserRepository
	Media MediaRepository
	S3    S3Repository
}

func Setup(connection db.Connection) *Repository {
	conn = connection
	return &Repository{
		User:  NewUserRepository(),
		Media: NewMediaRepository(),
		S3:    NewS3Repository(),
	}
}
