package repository

import (
	"backend/domain/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	User user.UserRepository
}

func NewRepositories(myDB *sqlx.DB) *Repositories {
	return &Repositories{
		User: NewUserRepository(myDB),
	}
}