package usecase

import (
	"backend/domain/user"
	"backend/pkg/errmsg"
	"backend/pkg/logutil"
)

type UserUsecase interface {
	GetByID(id int) (*user.User, error)
	GetAll() ([]user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id int) error
}

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

// GetByID retrieves a user by ID
func (u *userUsecase) GetByID(id int) (*user.User, error) {
	logutil.Info("Getting user by ID: %d", id)

	if id <= 0 {
		return nil, errmsg.NewError("ERR_BAD_REQUEST", "Invalid user ID")
	}

	user, err := u.userRepo.GetByID(id)
	if err != nil {
		if errmsg.GetErrorCode(err) == "ERR_NOT_FOUND" {
			return nil, err
		}
		return nil, errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to retrieve user", err)
	}
	return user, nil
}

// GetAll retrieves all users
func (u *userUsecase) GetAll() ([]user.User, error) {
	logutil.Info("Fetching all users")

	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to fetch users", err)
	}

	if len(users) == 0 {
		return nil, errmsg.NewError("ERR_NOT_FOUND", "No users found")
	}
	return users, nil
}

// Create inserts a new user
func (u *userUsecase) Create(user *user.User) error {
	logutil.Info("Creating user: %+v", user)

	if user.Name == "" || user.Email == "" {
		return errmsg.NewError("ERR_BAD_REQUEST", "Missing required fields: Name or Email")
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to create user", err)
	}
	return nil
}

// Update modifies an existing user
func (u *userUsecase) Update(user *user.User) error {
	logutil.Info("Updating user: %+v", user)

	if user.ID <= 0 {
		return errmsg.NewError("ERR_BAD_REQUEST", "Invalid user ID")
	}
	if user.Name == "" || user.Email == "" {
		return errmsg.NewError("ERR_BAD_REQUEST", "Missing required fields: Name or Email")
	}

	err := u.userRepo.Update(user)
	if err != nil {
		if errmsg.GetErrorCode(err) == "ERR_NOT_FOUND" {
			return err
		}
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to update user", err)
	}
	return nil
}

// Delete removes a user by ID
func (u *userUsecase) Delete(id int) error {
	logutil.Info("Deleting user with ID: %d", id)

	if id <= 0 {
		return errmsg.NewError("ERR_BAD_REQUEST", "Invalid user ID")
	}

	err := u.userRepo.Delete(id)
	if err != nil {
		if errmsg.GetErrorCode(err) == "ERR_NOT_FOUND" {
			return err
		}
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to delete user", err)
	}
	return nil
}
