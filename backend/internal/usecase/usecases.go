package usecase

import "backend/internal/repository"

type Usecases struct {
	User UserUsecase
}

func NewUsecases(repos *repository.Repositories) *Usecases {
	return &Usecases{
		User: NewUserUsecase(repos.User),
	}
}