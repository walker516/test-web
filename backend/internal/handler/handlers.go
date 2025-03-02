package handler

import "backend/internal/usecase"

type Handlers struct {
	User *UserHandler
}

func NewHandlers(usecase *usecase.Usecases) *Handlers {
	return &Handlers{
		User: NewUserHandler(usecase.User),
	}
}