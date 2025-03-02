package repository

import (
	"backend/domain/user"
	"backend/pkg/errmsg"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	*BaseSQLRepository
}

func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepository{NewBaseSQLRepository(db, filepath.Join("templates", "sql", "user_queries.sql"))}
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id int) (*user.User, error) {
	params := map[string]interface{}{"Id": id}
	var usr user.User

	err := r.ExecuteQuery("GetByID", params, func(rows *sqlx.Rows) error {
		if rows.Next() {
			if err := rows.StructScan(&usr); err != nil {
				return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to scan user data", err)
			}
			return nil
		}
		return errmsg.NewError("ERR_NOT_FOUND", "User not found")
	})

	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// GetAll retrieves all users
func (r *userRepository) GetAll() ([]user.User, error) {
	var users []user.User

	err := r.ExecuteQuery("GetAll", nil, func(rows *sqlx.Rows) error {
		for rows.Next() {
			var usr user.User
			if err := rows.StructScan(&usr); err != nil {
				return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to scan user data", err)
			}
			users = append(users, usr)
		}
		return nil
	})

	if err != nil {
		return nil, errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to fetch users", err)
	}

	return users, nil
}

// Create inserts a new user
func (r *userRepository) Create(user *user.User) error {
	params := map[string]interface{}{
		"Name":     user.Name,
		"Email":    user.Email,
		"Role":     user.Role,
	}

	_, err := r.ExecuteNonQuery("Create", params)
	if err != nil {
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to create user", err)
	}
	return nil
}

// Update modifies an existing user
func (r *userRepository) Update(user *user.User) error {
	params := map[string]interface{}{
		"Id":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
		"Role":  user.Role,
	}

	rowsAffected, err := r.ExecuteNonQuery("UpdateUser", params)
	if err != nil {
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to update user", err)
	}
	if rowsAffected == 0 {
		return errmsg.NewError("ERR_NOT_FOUND", "User not found")
	}
	return nil
}

// Delete removes a user by ID
func (r *userRepository) Delete(id int) error {
	params := map[string]interface{}{"Id": id}

	rowsAffected, err := r.ExecuteNonQuery("DeleteUser", params)
	if err != nil {
		return errmsg.WrapError("ERR_INTERNAL_SERVER", "Failed to delete user", err)
	}
	if rowsAffected == 0 {
		return errmsg.NewError("ERR_NOT_FOUND", "User not found")
	}
	return nil
}
